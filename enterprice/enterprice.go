package enterprice

import (
	"MINE/types"
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Enterprice struct {
	coalCapital atomic.Int64

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	mtx    sync.Mutex

	minerId       atomic.Int64
	activeMiners  map[int64]types.Miner
	minersHistory map[int64]types.Miner
}

func NewEterprice() *Enterprice {
	ctx, cancel := context.WithCancel(context.Background())

	return &Enterprice{
		ctx:           ctx,
		cancel:        cancel,
		activeMiners:  make(map[int64]types.Miner),
		minersHistory: make(map[int64]types.Miner),
	}
}

func (e *Enterprice) HireMiner(mHttp *types.HttpMiner) error {
	miners := map[string]func() types.Miner{
		"small":  types.NewSmallMiner,
		"normal": types.NewNormalMiner,
		"strong": types.NewStrongMiner}

	minerFunc, ok := miners[mHttp.MinerType]
	if !ok {
		return errors.New("Unknown miners type!")
	}

	m := minerFunc()
	totalPrice := int64(m.Info().Price * mHttp.Count)
	
	if e.coalCapital.Load() < totalPrice {
		return fmt.Errorf("Not enough coal! Current amount of coal is %d, but need %d", e.coalCapital.Load(), totalPrice)
	}
	
	e.coalCapital.Add(-totalPrice)

	if mHttp.Count > 1 {
		fmt.Printf("Шахтеры %s начали работу в количестве %d\n", m.Info().MinerType, mHttp.Count)
	} else {
		fmt.Printf("Шахтер %s начал работу\n", m.Info().MinerType)
	}

	for i := 0; i < mHttp.Count; i++ {
		m := minerFunc()
		e.addMiner(m)
	}

	go func() {
		e.wg.Wait()
		if mHttp.Count > 1 {
			fmt.Printf("Шахтеры %s закончили работу в количестве %d\n", m.Info().MinerType, mHttp.Count)
		} else {
			fmt.Printf("Шахтер %s закончил работу\n", m.Info().MinerType)
		}
	}()

	return nil
}

func (e *Enterprice) PassiveIncome() {
	for {
		select {
		case <-e.ctx.Done():
			fmt.Println("Game is stopped manualy!")
			return
		default:
			e.coalCapital.Add(1)
			time.Sleep(1 * time.Second)
		}

	}
}

func (e *Enterprice) addMiner(m types.Miner) {
	e.minerId.Add(1)
	currentID := e.minerId.Load()
	e.mtx.Lock()
	e.activeMiners[currentID] = m
	e.minersHistory[currentID] = m
	e.mtx.Unlock()

	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		coalCh := m.Run(e.ctx)
		for v := range coalCh {
			e.coalCapital.Add(int64(v))
		}
		e.mtx.Lock()
		delete(e.activeMiners, currentID)
		e.mtx.Unlock()
	}()

}

func (e *Enterprice) CapitalAmount() int64 {
	return e.coalCapital.Load()
}

func (e *Enterprice) Cancel() {
	e.cancel()
}

func (e *Enterprice) PrintActiveMiners() map[int64]types.Miner {
	return e.activeMiners
}

func (e *Enterprice) PrintHistoryMiners() map[int64]types.Miner {
	return e.minersHistory
}