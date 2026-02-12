package handlers

import (
	"MINE/enterprice"
	"io"
	"net/http"
)


func PurshareEquipmentHandle(w http.ResponseWriter, r *http.Request, e *enterprice.Enterprice) {
	if r.Method != http.MethodPost{
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	response, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Reading response error", http.StatusMethodNotAllowed)
		return
	}

	if err := e.PurshareEquipment(string(response)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))	
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully purchased a equipment!"))
}
