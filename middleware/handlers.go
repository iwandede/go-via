package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/iwandede/go-via/lib"
	"github.com/iwandede/go-via/models"
)

func HttpLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func CorsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Authtentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		tknStr := strings.Split(authorization, " ")
		if len(tknStr) < 1 {
			json.NewEncoder(w).Encode(lib.ResponseUnauthorized("Unauthorized!"))
			return
		}

		if tknStr[0] != "Bearer" {
			json.NewEncoder(w).Encode(lib.ResponseUnauthorized("Unauthorized!"))
			return
		}

		claims := &models.Claims{}

		tkn, err := jwt.ParseWithClaims(tknStr[1], claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("R5cCI6IkpXVCJ9"), nil
		})

		if err != nil && err == jwt.ErrSignatureInvalid {
			json.NewEncoder(w).Encode(lib.ResponseUnauthorized(err))
			return
		}

		if !tkn.Valid {
			json.NewEncoder(w).Encode(lib.ResponseUnauthorized("Unauthorized!"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
