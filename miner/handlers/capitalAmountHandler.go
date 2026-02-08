package handlers

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

func CapitalAmountHandler(w http.ResponseWriter, r *http.Request, coalCapital *atomic.Int64) {
	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Current amount of coal is %d", coalCapital.Load())
	w.Write([]byte(msg))
}
