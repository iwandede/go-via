package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/iwandede/go-via/config"

	"github.com/iwandede/go-via/lib"
	"github.com/jinzhu/gorm"
)

type Controllers struct {
	ctx       context.Context
	Datastore *gorm.DB
	Config    *config.Config
}

func NewControllers(ctx context.Context, config *config.Config, db *gorm.DB) *Controllers {
	return &Controllers{
		ctx:       ctx,
		Config:    config,
		Datastore: db,
	}
}

func (c Controllers) Index(w http.ResponseWriter, r *http.Request) {
	data := lib.ResponseSuccess("ping")
	json.NewEncoder(w).Encode(data)
}
