package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iwandede/go-via/controllers"
	"github.com/iwandede/go-via/middleware"
)

func (app APPServer) Routes(router *mux.Router) *mux.Router {
	c := controllers.NewControllers(app.ctx, app.Config, app.Datastore)

	router.HandleFunc("/auth", c.Signin).Methods("POST")
	router.HandleFunc("/ping", c.Index).Methods("GET")
	// users
	router.Handle("/users", middleware.Authtentication(http.HandlerFunc(c.GetAllUsers))).Methods("GET")
	router.Handle("/users/{id}", middleware.Authtentication(http.HandlerFunc(c.GetByID))).Methods("GET")
	router.Handle("/users/{id}", middleware.Authtentication(http.HandlerFunc(c.UpdateUsers))).Methods("PUT")
	router.Handle("/users/add", middleware.Authtentication(http.HandlerFunc(c.AddUsers))).Methods("POST")
	return router
}
