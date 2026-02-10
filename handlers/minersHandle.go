package handlers

import (
	"MINE/enterprice"
	"MINE/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func MinersHandle(w http.ResponseWriter, r *http.Request, e *enterprice.Enterprice) {
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	mHttp := types.HttpMiner{}
	if err := json.NewDecoder(io.Reader(r.Body)).Decode(&mHttp); err != nil {
		fmt.Println("err", err)
	}
	if mHttp.Count <= 0 {
		http.Error(w, "Number of miners cant be 0 or less", http.StatusBadRequest)
		return
	}

	err := e.HireMiner(&mHttp)
	if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))	
			return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully purchased a miner!"))
}
