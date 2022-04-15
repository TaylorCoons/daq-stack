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
	if e, ok := (err).(*NotAuthorized); ok {
		writeError(w, r, e, http.StatusUnauthorized)
	} else if e, ok := (err).(*MalformedBasicAuth); ok {
		writeError(w, r, e, http.StatusUnauthorized)
	} else if e, ok := (err).(auth.TokenNotAuthorized); ok {
		writeError(w, r, e, http.StatusUnauthorized)
	} else {
		panic(err)
	}
}

func IsAdminBasicAuthorized(f server.HandlerFunc) server.HandlerFunc {
	return server.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
		username, password, ok := r.BasicAuth()
		if !ok {
			handleAuthError(w, r, &MalformedBasicAuth{})
			return
		}
		if !auth.BasicAuth(username, password) {
			handleAuthError(w, r, &NotAuthorized{})
			return
		}
		f(ctx, w, r, p)
	})
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
	token, err := auth.CreateToken(c)
	if err != nil {
		handleAuthError(w, r, err)
		return
	}
	writeJson(w, token)
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
	writeJson(w, newToken)
}

func DeleteAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	err := auth.RevokeToken(c)
	if err != nil {
		handleAuthError(w, r, err)
		return
	}
}

func GetAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	writeJson(w, auth.SupportedAuthTypes())
}
