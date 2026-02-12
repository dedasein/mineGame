package handlers

import (
	"MINE/enterprice"
	"MINE/types"
	"fmt"
	"net/http"
)

func ActiveMinerHanlde(w http.ResponseWriter, r *http.Request, e *enterprice.Enterprice) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	sort := r.URL.Query().Get("sort")
	fmt.Println(sort)

	result := e.PrintActiveMiners()

	if len(result) == 0 {
		w.Write([]byte("There is no active miners"))
	} else {
		for id, miner := range result {
			info := miner.Info()
			if sort == "" || info.MinerType == sort {
				makeMessage(w, info, id)
			}
		}
	}
}

func makeMessage(w http.ResponseWriter, info types.MinerInfo, id int64) string {
	msg := fmt.Sprintf(
		"ID: %d | Type: %s | EnergyLeft: %d\n",
		id, info.MinerType, info.EnergyLeft)
	w.Write([]byte(msg))
	return msg
}
