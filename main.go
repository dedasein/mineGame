package main

import (
	"MINE/miner/handlers"
	"MINE/passiveIncome"
	"context"
	"net/http"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var coalCapital atomic.Int64 //общий баланс

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go passiveIncome.Increase(&coalCapital, ctx)

	http.HandleFunc("/capital", func(w http.ResponseWriter, r *http.Request) {
		handlers.CapitalAmountHandler(w, r, &coalCapital)
	})
	http.HandleFunc("/miners", func(w http.ResponseWriter, r *http.Request) {
		handlers.MinersHandle(ctx, w, r, &coalCapital, &wg)
	})
	http.HandleFunc("/cancel", func(w http.ResponseWriter, r *http.Request) {
		handlers.CancelHandler(cancel, w, r)
	})

	http.ListenAndServe(":9091", nil)
}
