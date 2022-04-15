package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TaylorCoons/daq-stack/src/connector"
	"github.com/TaylorCoons/daq-stack/src/models"
	"github.com/TaylorCoons/daq-stack/src/sdk/auth"
	server "github.com/TaylorCoons/gorouter"
)

func handleAuthError(w http.ResponseWriter, r *http.Request, err error) {
	if e, ok := (err).(*auth.NotAuthorized); ok {
		writeError(w, r, e, http.StatusUnauthorized)
	} else if e, ok := (err).(*MalformedBasicAuth); ok {
		writeError(w, r, e, http.StatusUnauthorized)
	} else {
		panic(err)
	}
}

func DevTest(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	username, password, ok := r.BasicAuth()
	if !ok {
		handleAuthError(w, r, &MalformedBasicAuth{})
		return
	}
	fmt.Printf("%s, %s", username, password)
}

func PostAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	username, password, ok := r.BasicAuth()
	if !ok {
		handleAuthError(w, r, &MalformedBasicAuth{})
		return
	}
	token, err := auth.CreateToken(c, username, password)
	if err != nil {
		handleAuthError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func PutAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	token := models.Token{}
	err := json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		handleAuthError(w, r, err)
		return
	}
	newToken, err := auth.RenewToken(c, token)
	if err != nil {
		handleAuthError(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newToken)
}

func DeleteAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	username, password, ok := r.BasicAuth()
	if !ok {
		handleAuthError(w, r, &MalformedBasicAuth{})
		return
	}
	err := auth.RevokeToken(c, username, password)
	if err != nil {
		handleAuthError(w, r, err)
		return
	}
}

func GetAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth.SupportedAuthTypes())
}
