package types

import (
	"context"
	"time"
)

type Miner interface {
	Run(ctx context.Context) <-chan int
	Info() MinerInfo
}

type MinerInfo struct {
	Price       int
	Energy      int
	Income      int
	BreakTime   int
	IncomeBonus int
	MinerType   string
}

type defaultMiner struct {
	price       int
	energy      int
	income      int
	breakTime   int
	incomeBonus int
	minerType   string
}

func (m *defaultMiner) Run(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= m.energy; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- m.income:
				time.Sleep(time.Duration(m.breakTime) * time.Second)
				if m.incomeBonus != 0 {
					m.income += m.incomeBonus
				}
			}
		}
	}()
	return ch
}

func (m defaultMiner) Info() MinerInfo {
	return MinerInfo{
		Price:       m.price,
		Energy:      m.energy,
		Income:      m.income,
		BreakTime:   m.breakTime,
		IncomeBonus: m.incomeBonus,
		MinerType:   m.minerType,
	}
}

func NewSmallMiner() Miner {
	return &defaultMiner{
		price:     5,
		energy:    30,
		income:    1,
		breakTime: 3,
		minerType: "small",
	}
}

func NewNormalMiner() Miner {
	return &defaultMiner{
		price:     50,
		energy:    45,
		income:    3,
		breakTime: 2,
		minerType: "normal",
	}
}

func NewStrongMiner() Miner {
	return &defaultMiner{
		price:       450,
		energy:      60,
		income:      10,
		breakTime:   1,
		incomeBonus: 3,
		minerType:   "strong",
	}
}
