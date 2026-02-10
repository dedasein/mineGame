package handlers

import (
	"MINE/enterprice"
	"fmt"
	"net/http"
)

func ActiveMinerHanlde(w http.ResponseWriter, r *http.Request, e *enterprice.Enterprice) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	result := e.PrintActiveMiners()

	if len(result) == 0 {
		w.Write([]byte("There is no active miners"))
	} else {
		for id, miner := range result {
			info := miner.Info()
			msg := fmt.Sprintf(
				"ID: %d | Type: %s | EnergyLeft: %d\n",
				id, info.MinerType, info.EnergyLeft)
			w.Write([]byte(msg))
		}

	}
}
