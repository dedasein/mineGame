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

	minerId            atomic.Int64
	activeMiners       map[int64]types.Miner
	minersHistory      map[int64]types.Miner
	pursharedEuqipment []*types.Equipment
}

func NewEterprice() *Enterprice {
	ctx, cancel := context.WithCancel(context.Background())

	return &Enterprice{
		ctx:                ctx,
		cancel:             cancel,
		activeMiners:       make(map[int64]types.Miner),
		minersHistory:      make(map[int64]types.Miner),
		pursharedEuqipment: make([]*types.Equipment, 0),
	}
}

// Нанять шахтера(ов)
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

	//Проверка достаточности угля
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

// Пассивный доход предприятия
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

// Добавления шахтера в историю
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

// Покупка оборудования
func (e *Enterprice) PurshareEquipment(s string, shutdownChan chan struct{}) error {
	equipment := []*types.Equipment{types.NewPickaxes(), types.NewVentilation(), types.NewCars()}

	//Проверка уникальности
	for i := range equipment {
		equipmentName, equipmentPrice := equipment[i].GetData()
		if s == equipmentName {
			for _, existingEquip := range e.pursharedEuqipment {
				existingName, _ := existingEquip.GetData()
				if existingName == equipmentName {
					return fmt.Errorf("Equipment '%s' has already been purchased", equipmentName)
				}
			}

			//Проверка достаточности угля для покупки
			if e.coalCapital.Load() >= int64(equipmentPrice) {
				e.coalCapital.Add(-int64(equipmentPrice))
				e.pursharedEuqipment = append(e.pursharedEuqipment, equipment[i])

				//Если длины одинаковые - завершаем игру
				if len(e.pursharedEuqipment) == len(equipment) {
					e.Cancel()
					fmt.Println("All the equipment has been purchased, and the enterprice has been shut down ")
					go func() {
		shutdownChan <- struct{}{}
	}()
					return nil
				}
				return nil
			} else {
				return fmt.Errorf("Not enough coal! Current amount of coal is %d, but need %d", e.coalCapital.Load(), equipmentPrice)
			}
		}
	}
	return fmt.Errorf("There is no such equipment. Available to purshare: ")
}

// Текущее кол-во угля
func (e *Enterprice) CapitalAmount() int64 {
	return e.coalCapital.Load()
}

// Ручная остановка
func (e *Enterprice) Cancel() {
	e.cancel()
}

// Запущенные в данный момент шахтеры
func (e *Enterprice) PrintActiveMiners() map[int64]types.Miner {
	return e.activeMiners
}

// Все когда-либо приобретенные шахтеры
func (e *Enterprice) PrintHistoryMiners() map[int64]types.Miner {
	return e.minersHistory
}

// УДАЛИТЬ
func (e *Enterprice) CheatCoal() {
	e.coalCapital.Store(1_000_000)
}
