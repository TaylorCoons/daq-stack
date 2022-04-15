package handlers

import (
	"errors"
	"fmt"
	"net/http"

	server "github.com/TaylorCoons/gorouter"
)

func Middleware(w http.ResponseWriter, r *http.Request, p server.PathParams, h server.HandlerFunc) {
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println(rec)
			writeError(w, r, errors.New("something went wrong 😕"), http.StatusInternalServerError)
		}
	}()
	h(nil, w, r, p)
}
