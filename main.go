package main

import (
	"MINE/enterprice"
	"MINE/handlers"
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	e := enterprice.NewEterprice()
	go e.PassiveIncome()

	// Канал для сигнала остановки сервера
	shutdownChan := make(chan struct{})

	mux := http.NewServeMux()

	mux.HandleFunc("/miners", func(w http.ResponseWriter, r *http.Request) {
		handlers.MinersHandle(w, r, e)
	})

	mux.HandleFunc("/capital", func(w http.ResponseWriter, r *http.Request) {
		handlers.CapitalHandler(w, r, e)
	})

	mux.HandleFunc("/miners/active", func(w http.ResponseWriter, r *http.Request) {
		handlers.ActiveMinerHanlde(w, r, e)
	})

	mux.HandleFunc("/miners/history", func(w http.ResponseWriter, r *http.Request) {
		handlers.HistoryMinerHanlde(w, r, e)
	})

	mux.HandleFunc("/equipment", func(w http.ResponseWriter, r *http.Request) {
		handlers.PurshareEquipmentHandle(w, r, e, shutdownChan)
	})

	mux.HandleFunc("/miners/price", handlers.MinersPriceHanlde)

	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		handlers.CancelHandler(w, r, e, shutdownChan)
	})

	//УДАЛИТЬ
	mux.HandleFunc("/cheat", func(w http.ResponseWriter, r *http.Request) {
		handlers.CheatHandler(w, r, e)
	})

	server := &http.Server{
		Addr:    ":9091",
		Handler: mux,
	}
	
	// Запускаем сервер в отдельной горутине
	go func() {
		log.Println("Server started on :9091")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	// Ждём сигнала на остановку
	<-shutdownChan

	// Корректно останавливаем сервер с timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v\n", err)
	} else {
		log.Println("Server has been gracefully shut down")
	}
}
