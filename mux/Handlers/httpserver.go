package Handlers

import (
	"main-api-store-management/mux/muxlib"
	"net/http"
)

func NotFoundController(w http.ResponseWriter, r *http.Request) {
	muxlib.HttpError404(w, "Requested resource doesn't exist. Please check your path.")
}
