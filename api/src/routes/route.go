package routes

import (
	"github.com/TaylorCoons/daq-stack/src/handlers"
	server "github.com/TaylorCoons/gorouter"
)

var Routes []server.Route = []server.Route{
	{Method: "GET", Path: "/health", Handler: handlers.GetHealth},

	{Method: "GET", Path: "/auth", Handler: handlers.GetAuth},
	{Method: "POST", Path: "/auth", Handler: handlers.IsAdminBasicAuthorized(handlers.PostAuth)},
	{Method: "PUT", Path: "/auth", Handler: handlers.IsAdminTokenAuthorized(handlers.PutAuth)},
	{Method: "DELETE", Path: "/auth", Handler: handlers.IsAdminBasicAuthorized(handlers.DeleteAuth)},

	{Method: "POST", Path: "/devtest", Handler: handlers.DevTest},
}
