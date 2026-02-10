package types

import (
	"context"
	"sync/atomic"
	"time"
)

type Miner interface {
	Run(ctx context.Context) <-chan int
	Info() MinerInfo
}

type MinerInfo struct {
	Price       int
	EnergyLeft  int
	MinerType   string
}

type defaultMiner struct {
	price       int
	energy      atomic.Int64
	income      int
	breakTime   int
	incomeBonus int
	minerType   string
}

func (m *defaultMiner) Run(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for m.energy.Load() > 0 {
			select {
			case <-ctx.Done():
				return
			case ch <- m.income:
				time.Sleep(time.Duration(m.breakTime) * time.Second)
				if m.incomeBonus != 0 {
					m.income += m.incomeBonus
				}
				m.energy.Add(-1)
			}
		}
	}()
	return ch
}

func (m *defaultMiner) Info() MinerInfo {
	return MinerInfo{
		Price:       m.price,
		EnergyLeft:  int(m.energy.Load()),
		MinerType:   m.minerType,
	}
}

func NewSmallMiner() Miner {
	m := &defaultMiner{
		price:     5,
		income:    1,
		breakTime: 3,
		minerType: "small",
	}
	m.energy.Store(30)
	return m
}

func NewNormalMiner() Miner {
	m := &defaultMiner{
		price:     50,
		income:    3,
		breakTime: 2,
		minerType: "normal",
	}
	m.energy.Store(45)
	return m
}

func NewStrongMiner() Miner {
	m := &defaultMiner{
		price:       450,
		income:      10,
		breakTime:   1,
		incomeBonus: 3,
		minerType:   "strong",
	}
	m.energy.Store(60)
	return m
}
