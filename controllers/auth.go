package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iwandede/go-via/lib"
	"github.com/iwandede/go-via/models"
)

func (Users Controllers) Signin(w http.ResponseWriter, r *http.Request) {
	var Creds models.Credentials
	var Response models.Users

	if err := json.NewDecoder(r.Body).Decode(&Creds); err != nil {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest(err))
		return
	}

	Auth := Users.Datastore.Where("username = ? OR email = ?", Creds.Email, Creds.Email).Find(&Response).Order("created_at DESC").Limit(1)
	if Auth.Error != nil {
		json.NewEncoder(w).Encode(lib.ResponseInternalError(Auth.Error))
		return
	}

	ok := lib.CheckPasswordHash(Creds.Password, Response.Password)
	if !ok {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest("Password not match!"))
		return
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	claims := &models.Claims{
		Username: Creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Users.Config.Security.SecretKey))
	if err != nil {
		json.NewEncoder(w).Encode(lib.ResponseBadRequest("Failed sign token"))
		return
	}

	responseData := &models.ResponseAuth{
		ID:    Response.ID,
		Email: Response.Email,
		Token: tokenString,
	}

	json.NewEncoder(w).Encode(lib.ResponseSuccess(responseData))
	return
}
