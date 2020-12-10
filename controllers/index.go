package controllers

import (
	"context"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"net/http"
	"github.com/iwandede/go-via/config"
	"github.com/iwandede/go-via/lib"
)

type Controllers struct {
	ctx       		context.Context
	Datastore 		*sqlx.DB
	Config    		*config.Config
}

func NewControllers(ctx context.Context, config *config.Config, db *sqlx.DB) *Controllers {
	return &Controllers{
		ctx:       ctx,
		Config:    config,
		Datastore: db,
	}
}

func (c Controllers) Index(w http.ResponseWriter, r *http.Request) {
	data := lib.ResponseSuccess("API Version V1.0.0", )
	json.NewEncoder(w).Encode(data)
}

func (c Controllers) Ping(w http.ResponseWriter, r *http.Request) {
	data := lib.ResponseSuccess("pong")
	json.NewEncoder(w).Encode(data)
}