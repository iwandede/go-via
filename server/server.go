package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iwandede/go-via/config"
	"github.com/iwandede/go-via/database"
	"github.com/iwandede/go-via/lib"
	"github.com/iwandede/go-via/middleware"
	"github.com/jinzhu/gorm"
)

type APPServer struct {
	Config    *config.Config
	Server    *config.Server
	Datastore *gorm.DB
	ctx       context.Context
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
	router := mux.NewRouter()
	router.Use(middleware.CorsHeaders)
	router.Use(middleware.HttpLogging)

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

	return app.Routes(router)
}
