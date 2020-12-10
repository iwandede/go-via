package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"

	"github.com/iwandede/go-via/lib"
	"github.com/iwandede/go-via/models"
)

func (Service Controllers) GetAllService(w http.ResponseWriter, r *http.Request) {
	var ServiceDTO []*models.Service
	query := "SELECT * FROM  workbench.app_service ORDER BY srv_created_at DESC"
	if err := Service.Datastore.Select(&ServiceDTO, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			json.NewEncoder(w).Encode(lib.ResponseBadRequest("No file!"))
			return
		}
		json.NewEncoder(w).Encode(lib.ResponseBadRequest(err))
		return
	}

	json.NewEncoder(w).Encode(lib.ResponseSuccess(ServiceDTO))
	return
}

func (Service Controllers) AddService(w http.ResponseWriter, r *http.Request) {
	var ServiceDTO models.Service
	var t time.Time

	if err := json.NewDecoder(r.Body).Decode(&ServiceDTO); err != nil {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest(err))
		return
	}

	PrivateKey, err := lib.GenerateSalt(lib.GenerateID())
	if err != nil {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest(err))
		return
	}

	ServiceDTO.ID = uuid.New()
	ServiceDTO.Signature = lib.EncodeHMACSHA256(fmt.Sprintf("%v", ServiceDTO.ID), PrivateKey)
	ServiceDTO.PrivateKey = PrivateKey
	ServiceDTO.Status = 1

	cols := []string{"srv_id", "srv_name", "srv_description", "srv_signature", "srv_private_key", "srv_status", "srv_created_at", "srv_updated_at"}
	values := []string{ServiceDTO.ID.String(), ServiceDTO.Name, ServiceDTO.Description, ServiceDTO.Signature, ServiceDTO.PrivateKey, lib.ToString(ServiceDTO.Status), t.Format("2006-01-02 15:04:05"), t.Format("2006-01-02 15:04:05")}
	query := fmt.Sprintf("INSERT INTO workbench.app_service (%v) VALUES ('%v') RETURNING *", strings.Join(cols, ","), strings.Join(values, "','"))

	if err := Service.Datastore.QueryRowx(query).StructScan(&ServiceDTO); err != nil {
		json.NewEncoder(w).Encode(lib.ResponseInternalError(err))
		return
	}

	json.NewEncoder(w).Encode(lib.ResponseSuccess(ServiceDTO))
	return
}
