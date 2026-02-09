package main

import (
	"MINE/enterprice"
	"MINE/handlers"
	"net/http"
)

func main() {
	e := enterprice.NewEterprice()
	go e.PassiveIncome()

	http.HandleFunc("/miners", func(w http.ResponseWriter, r *http.Request) {
		handlers.MinersHandle(w, r, e)
	})

	http.HandleFunc("/capital", func(w http.ResponseWriter, r *http.Request) {
		handlers.CapitalAmountHandler(w, r, e)
	})

	http.ListenAndServe(":9091", nil)
}
