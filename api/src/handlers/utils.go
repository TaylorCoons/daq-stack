package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TaylorCoons/daq-stack/src/models"
)

func writeJson(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, r *http.Request, err error, status int) {
	w.WriteHeader(status)
	httpErr := models.HttpError{
		Timestamp: time.Now().UTC().String(),
		Error:     http.StatusText(status),
		Status:    status,
		Message:   err.Error(),
		Path:      r.URL.Path,
	}
	writeJson(w, httpErr)
}
