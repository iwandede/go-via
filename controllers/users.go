package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iwandede/go-via/lib"
	"github.com/iwandede/go-via/models"
)

func (Users Controllers) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var Response []models.Users
	var Limit, Offset int64
	qs := r.URL.Query()

	if qs.Get("from") != "" || qs.Get("limit") != "" {
		Limit, Offset = 10, 0
		qs.Del("from")
		qs.Del("limit")
	}

	Params := make(map[string]interface{})
	for field, values := range qs {
		Params[field] = values[0]
	}
	//pagination := ""
	user := Users.Datastore.Where(Params).Limit(Limit).Offset(Offset).Find(&Response).Order("created_at DESC")
	json.NewEncoder(w).Encode(lib.ResponseSuccess(user.Value))
	return
}

func (Users Controllers) GetByID(w http.ResponseWriter, r *http.Request) {
	var Response models.Users
	params := mux.Vars(r)
	user := Users.Datastore.Where("id = ?", params["id"]).Find(&Response)

	json.NewEncoder(w).Encode(lib.ResponseSuccess(user.Value))
	return
}

func (Users Controllers) AddUsers(w http.ResponseWriter, r *http.Request) {
	var UsersDTO models.Users

	if err := json.NewDecoder(r.Body).Decode(&UsersDTO); err != nil {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest(err))
		return
	}

	password, err := lib.GeneratePassword(UsersDTO.Password)
	if err != nil {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest(err))
		return
	}

	UsersDTO.Password = password
	if err := Users.Datastore.Create(&UsersDTO).Error; err != nil {
		json.NewEncoder(w).Encode(lib.ResponseInternalError(err))
		return
	}

	//defer Users.Datastore.Close()
	json.NewEncoder(w).Encode(lib.ResponseSuccess(UsersDTO))
	return
}

func (Users Controllers) UpdateUsers(w http.ResponseWriter, r *http.Request) {
	var UsersDTO models.Users
	params := mux.Vars(r)
	UserID := params["id"]
	if err := json.NewDecoder(r.Body).Decode(&UsersDTO); err != nil {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest(err))
		return
	}

	update := Users.Datastore.Model(&UsersDTO).Where("id = ?", UserID).Updates(&UsersDTO).Last(&UsersDTO)
	if update.Error != nil {
		json.NewEncoder(w).Encode(lib.ResponseInternalError(update.Error))
		return
	}

	if update.RowsAffected < 1 {
		errorMessage := fmt.Sprintf("UsersID %v not found in record!", UserID)
		json.NewEncoder(w).Encode(lib.ResponseNotFound(errorMessage))
		return
	}

	//defer Users.Datastore.Close()
	json.NewEncoder(w).Encode(lib.ResponseSuccess(update.Value))
	return
}
