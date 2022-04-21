package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/TaylorCoons/daq-stack/src/connector"
	"github.com/TaylorCoons/daq-stack/src/models"
	"github.com/TaylorCoons/daq-stack/src/sdk/app"
	server "github.com/TaylorCoons/gorouter"
)

func handleAppError(w http.ResponseWriter, r *http.Request, err error) {
	if e, ok := (err).(app.AppNotFoundError); ok {
		writeError(w, r, e, http.StatusNotFound)
	} else {
		panic(err)
	}

}

func PostCreateApp(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	a := models.App{}
	json.NewDecoder(r.Body).Decode(&a)
	created, err := app.CreateApp(c, a)
	if err != nil {
		handleAppError(w, r, err)
		return
	}
	writeJson(w, created, http.StatusCreated)
}

func PutUpdateApp(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	a := models.App{}
	json.NewDecoder(r.Body).Decode(&a)
	id := p["appId"]
	updated, err := app.UpdateApp(c, a, id)
	if err != nil {
		handleAppError(w, r, err)
		return
	}
	writeJson(w, updated, http.StatusOK)
}

func GetApp(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	a, err := app.GetApp(c, p["appId"])
	if err != nil {
		handleAppError(w, r, err)
		return
	}
	writeJson(w, a, http.StatusOK)
}

func DeleteApp(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	err := app.DeleteApp(c, p["appId"])
	if err != nil {
		handleAppError(w, r, err)
		return
	}
	writeStatus(w, http.StatusOK)
}

func GetListApps(ctx context.Context, w http.ResponseWriter, r *http.Request, p server.PathParams) {
	c := connector.Get()
	res, err := app.ListApps(c)
	if err != nil {
		handleAppError(w, r, err)
		return
	}
	writeJson(w, res, http.StatusOK)
}
