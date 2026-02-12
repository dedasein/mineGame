package handlers

import (
	"MINE/enterprice"
	"net/http"
)

func CheatHandler(w http.ResponseWriter, r *http.Request, e *enterprice.Enterprice) {
	e.CheatCoal()
}

//УДАЛИТЬ
