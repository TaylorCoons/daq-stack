package handlers

import (
	"errors"
	"net/http"

	server "github.com/TaylorCoons/gorouter"
)

func Middleware(w http.ResponseWriter, r *http.Request, p server.PathParams, h server.HandlerFunc) {
	defer func() {
		if rec := recover(); rec != nil {
			writeError(w, r, errors.New("something went wrong 😕"), http.StatusInternalServerError)
		}
	}()
	h(nil, w, r, p)
}
