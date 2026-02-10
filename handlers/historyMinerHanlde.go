package handlers

import (
	"MINE/enterprice"
	"fmt"
	"net/http"
)

func HistoryMinerHanlde(w http.ResponseWriter, r *http.Request, e *enterprice.Enterprice) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	result := e.PrintHistoryMiners()

	if len(result) == 0 {
		w.Write([]byte("History is empty"))
	} else {
		smallMinerCounter := 0
		normalMinerCounter := 0
		strongMinerCounter := 0
		for _, miner := range result {
			switch miner.Info().MinerType {
			case "small":
				smallMinerCounter++
			case "normal":
				normalMinerCounter++
			case "strong":
				strongMinerCounter++
			}
		}
		msg := fmt.Sprintf("Total of small miners ever worked %d\nTotal of normal miners ever worked %d\nTotal of strong miners ever worked %d\n",
			smallMinerCounter, normalMinerCounter, strongMinerCounter)

		w.Write([]byte(msg))
	}
}
