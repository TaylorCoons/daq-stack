package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TaylorCoons/daq-stack/src/connector"
	"github.com/TaylorCoons/daq-stack/src/sdk/auth"
	server "github.com/TaylorCoons/gorouter"
)

func handleAuthError(w http.ResponseWriter, r *http.Request, err error) {
	if e, ok := (err).(*NotAuthorized); ok {
		writeError(w, r, e, http.StatusUnauthorized)
	} else if e, ok := (err).(*MalformedBasicAuth); ok {
		writeError(w, r, e, http.StatusUnauthorized)
	} else if e, ok := (err).(*TokenNotAuthorized); ok {
		writeError(w, r, e, http.StatusUnauthorized)
	} else if e, ok := (err).(*NoApiKeyProvided); ok {
		writeError(w, r, e, http.StatusUnauthorized)
	} else {
		panic(err)
	}
}

func IsAdminBasicAuthorized(f server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
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
	}
}

func IsAdminTokenAuthorized(f server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
		client := connector.Get()
		key := r.Header.Get("x-api-key")
		if len(key) == 0 {
			handleAuthError(w, r, &NoApiKeyProvided{})
			return
		}
		if !auth.ValidateToken(client, key) {
			handleAuthError(w, r, &TokenNotAuthorized{})
			return
		}
		f(ctx, w, r, p)
	}
}

func DevTest(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	fmt.Println("test")
	fmt.Println(r.Header.Get("x-api-key"))
}

func PostAuthLogin(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	token, err := auth.CreateToken(c)
	if err != nil {
		handleAuthError(w, r, err)
		return
	}
	writeJson(w, token)
}

func PutAuthRenew(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	newToken, err := auth.RenewToken(c, r.Header.Get("x-api-key"))
	if err != nil {
		handleAuthError(w, r, err)
		return
	}
	writeJson(w, newToken)
}

func DeleteAuthRelease(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	err := auth.RevokeToken(c, r.Header.Get("x-api-key"))
	if err != nil {
		handleAuthError(w, r, err)
		return
	}
}

func DeleteAuthRevoke(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	err := auth.RevokeAll(c)
	if err != nil {
		handleAuthError(w, r, err)
		return
	}
}

func GetAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	writeJson(w, auth.SupportedAuthTypes())
}
