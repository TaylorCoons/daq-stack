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

// TODO: REMOVE ME:
var encode string = "YWRtaW46cGFzcw=="

func PostAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	//YWRtaW46cGFzcw==
	token, err := auth.CreateToken(c, encode)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

func PutAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	token := models.Token{}
	err := json.NewDecoder(r.Body).Decode(&token)
	fmt.Println(token)
	if err != nil {
		panic(err)
	}
	newToken, err := auth.RenewToken(c, token)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newToken)
}

func DeleteAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	err := auth.RevokeToken(c, encode)
	if err != nil {
		panic(err)
	}
}

func GetAuth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(auth.SupportedAuthTypes())
}
