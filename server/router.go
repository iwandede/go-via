package server

import (
	"github.com/gorilla/mux"
	"github.com/iwandede/go-via/controllers"
	"github.com/iwandede/go-via/middleware"
	"net/http"
)

func (app APPServer) Routes(mw *middleware.ConfigMiddleware, router *mux.Router) *mux.Router {
	C := controllers.NewControllers(app.ctx, app.Config, app.Datastore)
	router.HandleFunc("/", C.Index).Methods("GET")
	router.HandleFunc("/ping", C.Ping).Methods("GET")
	// Service
	router.Handle("/service", mw.AuthenticationGuard(http.HandlerFunc(C.GetAllService))).Methods("GET")
	router.Handle("/service/add", mw.AuthenticationGuard(http.HandlerFunc(C.AddService))).Methods("POST")
	// Notification
	router.Handle("/send-notification/whatsapp", mw.AuthenticationGuard(http.HandlerFunc(C.SendWhatsApp))).Methods("POST")
	router.Handle("/send-notification/sms", mw.AuthenticationGuard(http.HandlerFunc(C.SendSMS))).Methods("POST")

	return router
}
