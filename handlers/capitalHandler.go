package handlers

import (
	"MINE/enterprice"
	"fmt"
	"net/http"
)

func CapitalHandler(w http.ResponseWriter, r *http.Request, e *enterprice.Enterprice) {
	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Current amount of coal is %d", e.CapitalAmount())
	w.Write([]byte(msg))
}
