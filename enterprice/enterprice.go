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

	ctx context.Context
	cancel context.CancelFunc
	wg sync.WaitGroup
}

func NewEterprice() *Enterprice{
	ctx, cancel := context.WithCancel(context.Background())

	return &Enterprice{
		ctx: ctx,
		cancel: cancel,
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
	if e.coalCapital.Load() < int64(m.Info().Price*mHttp.Count) {
		return errors.New(fmt.Sprintf("Not enough coal! Current amount of coal is %d", e.coalCapital.Load()))
	} else {
		e.coalCapital.Add(-int64(m.Info().Price*mHttp.Count))
	}

	if mHttp.Count > 1 {
		fmt.Printf("Шахтеры %s начали работу в количестве %d\n", m.Info().MinerType, mHttp.Count)
	} else {
		fmt.Printf("Шахтер %s начал работу\n", m.Info().MinerType)
	}

	for i := 0; i < mHttp.Count; i++ {
		e.wg.Add(1)
		go func() {
			defer e.wg.Done()
			coal := m.Run(e.ctx)
			for v := range coal {
				e.coalCapital.Add(int64(v))
			}
		}()
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


func (e *Enterprice) CapitalAmount() int64{
	return e.coalCapital.Load()
}

func (e *Enterprice) Cancel() {
	e.cancel()
}

