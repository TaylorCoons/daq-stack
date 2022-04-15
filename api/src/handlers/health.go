package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	server "github.com/TaylorCoons/gorouter"
)

func GetHealth(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}
