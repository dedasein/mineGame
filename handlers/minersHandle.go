package handlers

import (
	"MINE/miner"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
)

type httpMiner struct {
	MinerType string `json:"type"`
	Count     int    `json:"count"`
}

func MinersHandle(ctx context.Context, w http.ResponseWriter, r *http.Request, coalCapital *atomic.Int64, wg *sync.WaitGroup) {
	mHttp := httpMiner{}
	if err := json.NewDecoder(io.Reader(r.Body)).Decode(&mHttp); err != nil {
		fmt.Println("err", err)
	}

	if mHttp.Count <= 0 {
		http.Error(w, "Number of miners cant be 0 or less", http.StatusBadRequest)
		return
	}

	miners := map[string]func() miner.Miner{"small": miner.NewSmallMiner, "normal": miner.NewNormalMiner, "strong": miner.NewStrongMiner}

	minerFunc, ok := miners[mHttp.MinerType]
	if !ok {
		http.Error(w, "Unknown miner type", http.StatusBadRequest)
		return
	}

	m := minerFunc()
	if int(coalCapital.Load()) < m.Info().Price*mHttp.Count {
		msg := fmt.Sprintf("Not enough coal! Current amount of coal is %d", coalCapital.Load())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if mHttp.Count > 1 {
		fmt.Printf("Шахтеры %s начали работу в количестве %d\n", m.Info().MinerType, mHttp.Count)
	} else {
		fmt.Printf("Шахтер %s начал работу\n", m.Info().MinerType)
	}

	for i := 0; i < mHttp.Count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			coal := m.Run(ctx)
			for v := range coal {
				coalCapital.Add(int64(v))
			}
		}()
	}

	go func() {
		wg.Wait()
		if mHttp.Count > 1 {
			fmt.Printf("Шахтеры %s закончили работу в количестве %d\n", m.Info().MinerType, mHttp.Count)
		} else {
			fmt.Printf("Шахтер %s закончил работу\n", m.Info().MinerType)
		}
	}()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully purchased a miner!"))
}
