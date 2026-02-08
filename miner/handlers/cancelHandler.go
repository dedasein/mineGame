package handlers

import (
	"context"
	"net/http"
)

func CancelHandler(cancel context.CancelFunc, w http.ResponseWriter, r *http.Request) {
	cancel()
}

//TODO write responce
