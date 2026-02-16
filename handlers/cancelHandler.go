package handlers

import (
	"MINE/enterprice"
	"net/http"
)

func CancelHandler(w http.ResponseWriter, r *http.Request, e *enterprice.Enterprice, shutdownChan chan struct{}) {
	// Останавливаем бизнес-логику
	e.Cancel()

	// Отправляем ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"Server is shutting down gracefully"}`))

	// Отправляем сигнал на остановку сервера
	go func() {
		shutdownChan <- struct{}{}
	}()
}
