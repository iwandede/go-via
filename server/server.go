package server

import (
	"context"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iwandede/go-via/config"
	"github.com/iwandede/go-via/database"
	"github.com/iwandede/go-via/lib"
	"github.com/iwandede/go-via/middleware"
)

type APPServer struct {
	Config      *config.Config
	Server      *config.Server
	Datastore   *sqlx.DB
	ctx         context.Context
}

func NewAppHttpServer(config *config.Config) *APPServer {
	ds, err := database.DataStore(config)
	if err != nil {
		log.Fatalf("failed to connect database %v", err)
	}
	return &APPServer{
		Config:    config,
		Server:    config.Server,
		Datastore: ds,
	}
}

func (app APPServer) InitRouter() *mux.Router {
	middlewares := middleware.NewMiddlewareConfig(app.Config.Security, app.Datastore)
	router := mux.NewRouter()
	router.Use(middlewares.CorsHeaders)
	router.Use(middlewares.HttpLogging)
	// router.Use(middleware.Authtentication)
	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(lib.ResponseMethodNotAllowed("405 Method not allowed"))
	})
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(lib.ResponseNotFound("404 Page not found!"))
	})

	return app.Routes(middlewares, router)
}
