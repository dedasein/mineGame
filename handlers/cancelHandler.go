package handlers

import (
	"MINE/enterprice"
	"net/http"
)

func CancelHandler(w http.ResponseWriter, r *http.Request, e *enterprice.Enterprice) {
	e.Cancel()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Enterprice has been successfully shut down."))
}

//TODO write responce
