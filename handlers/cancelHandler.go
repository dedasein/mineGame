package handlers

import (
	"MINE/enterprice"
	"context"
	"net/http"
)

func CancelHandler(cancel context.CancelFunc, w http.ResponseWriter, e *enterprice.Enterprice) {
	e.Cancel()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Enterprice has been successfully shut down."))
}

//TODO write responce
