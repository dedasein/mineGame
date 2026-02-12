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
		handlers.CapitalHandler(w, r, e)
	})

	http.HandleFunc("/miners/active", func(w http.ResponseWriter, r *http.Request) {
		handlers.ActiveMinerHanlde(w, r, e)
	})

	http.HandleFunc("/miners/history", func(w http.ResponseWriter, r *http.Request) {
		handlers.HistoryMinerHanlde(w, r, e)
	})

	http.HandleFunc("/equipment", func(w http.ResponseWriter, r *http.Request) {
		handlers.PurshareEquipmentHandle(w, r, e)
	})

	http.HandleFunc("/miners/price", handlers.MinersPriceHanlde)

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request){
		handlers.CancelHandler(w, r, e)
	})

	//УДАЛИТЬ
	http.HandleFunc("/cheat", func(w http.ResponseWriter, r *http.Request) {
		handlers.CheatHandler(w, r, e)
	})

	http.ListenAndServe(":9091", nil)
}
