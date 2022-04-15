package routes

import (
	"github.com/TaylorCoons/daq-stack/src/handlers"
	server "github.com/TaylorCoons/gorouter"
)

var Routes []server.Route = []server.Route{
	{Method: "GET", Path: "/health", Handler: handlers.GetHealth},

	{Method: "POST", Path: "/auth", Handler: handlers.PostAuth},
	{Method: "GET", Path: "/auth", Handler: handlers.GetAuth},
	{Method: "DELETE", Path: "/auth", Handler: handlers.DeleteAuth},
}
