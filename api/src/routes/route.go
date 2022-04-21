package routes

import (
	"github.com/TaylorCoons/daq-stack/src/handlers"
	server "github.com/TaylorCoons/gorouter"
)

var Routes []server.Route = []server.Route{
	{Method: "GET", Path: "/health", Handler: handlers.GetHealth},

	{Method: "GET", Path: "/auth", Handler: handlers.GetAuth},
	{Method: "POST", Path: "/auth/login", Handler: handlers.IsAdminBasicAuthorized(handlers.PostAuthLogin)},
	{Method: "PUT", Path: "/auth/renew", Handler: handlers.IsAdminTokenAuthorized(handlers.PutAuthRenew)},
	{Method: "DELETE", Path: "/auth/release", Handler: handlers.IsAdminTokenAuthorized(handlers.DeleteAuthRelease)},
	{Method: "DELETE", Path: "/auth/revoke", Handler: handlers.IsAdminBasicAuthorized(handlers.DeleteAuthRevoke)},

	{Method: "GET", Path: "/a", Handler: handlers.IsAdminTokenAuthorized(handlers.GetListApps)},
	{Method: "POST", Path: "/a", Handler: handlers.IsAdminTokenAuthorized(handlers.PostCreateApp)},
	{Method: "GET", Path: "/a/:appId", Handler: handlers.IsAdminTokenAuthorized(handlers.GetApp)},
	{Method: "PUT", Path: "/a/:appId", Handler: handlers.IsAdminTokenAuthorized(handlers.PutUpdateApp)},
	{Method: "DELETE", Path: "/a/:appId", Handler: handlers.IsAdminTokenAuthorized(handlers.DeleteApp)},

	{Method: "POST", Path: "/devtest", Handler: handlers.DevTest},
}
