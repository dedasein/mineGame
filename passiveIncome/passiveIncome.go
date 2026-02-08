package passiveIncome

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

func Increase(coalCapital *atomic.Int64, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Game is stopped manualy!")
			return
		default:
			coalCapital.Add(1)
			time.Sleep(1 * time.Second)
		}

	}
}
