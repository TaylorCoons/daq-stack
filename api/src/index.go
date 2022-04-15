package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/TaylorCoons/daq-stack/src/connector"
	"github.com/TaylorCoons/daq-stack/src/handlers"
	"github.com/TaylorCoons/daq-stack/src/routes"
	"github.com/TaylorCoons/daq-stack/src/sdk/auth"
	server "github.com/TaylorCoons/gorouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	startServer()
}

func startServer() {

	c, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(fmt.Errorf("failed to create mongo client: %v", err))
	}
	connector.Set(c)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = c.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer c.Disconnect(ctx)

	auth.IndexTables(c)

	compiledRoutes := server.CompileRoutes(routes.Routes)
	server := server.Server{CompiledRoutes: compiledRoutes}
	server.Middleware = handlers.Middleware

	port := 8080
	bind := fmt.Sprintf(":%d", port)
	fmt.Printf("Parrot is listening on: %d\n", port)

	err = http.ListenAndServe(bind, server)
	if err != nil {
		panic(err)
	}
}
