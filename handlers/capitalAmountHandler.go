package handlers

import (
	"MINE/enterprice"
	"fmt"
	"net/http"
)

func CapitalAmountHandler(w http.ResponseWriter, r *http.Request, e *enterprice.Enterprice) {
	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Current amount of coal is %d", e.CapitalAmount())
	w.Write([]byte(msg))
}
