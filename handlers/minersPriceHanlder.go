package handlers

import (
	"MINE/types"
	"fmt"
	"net/http"
)

func MinersPriceHanlde(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	miners := []types.Miner{types.NewSmallMiner(), types.NewNormalMiner(), types.NewStrongMiner()}
	for i := range miners {
		msg := fmt.Sprintf("Type: %s | Price: %v\n", miners[i].Info().MinerType, miners[i].Info().Price)
		w.Write([]byte(msg))
	}
}