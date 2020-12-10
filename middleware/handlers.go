package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"

	"github.com/iwandede/go-via/models"

	"github.com/iwandede/go-via/config"
	"github.com/iwandede/go-via/lib"
)

type ConfigMiddleware struct {
	key       *config.Security
	Datastore *sqlx.DB
}

func NewMiddlewareConfig(key *config.Security, db *sqlx.DB) *ConfigMiddleware {
	return &ConfigMiddleware{
		key:       key,
		Datastore: db,
	}
}

func (m ConfigMiddleware) HttpLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (m ConfigMiddleware) CorsHeaders(next http.Handler) http.Handler {
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

func (m ConfigMiddleware) AuthenticationGuard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data = models.Service{}
		token := r.Header.Get("token")
		signature := r.Header.Get("signature")
		query := fmt.Sprintf("SELECT * FROM workbench.app_service WHERE srv_id = '%s' LIMIT 1", token)

		if token == "" {
			json.NewEncoder(w).Encode(lib.ResponseBadRequest("Invalid Token"))
			return
		}

		if signature == "" {
			json.NewEncoder(w).Encode(lib.ResponseBadRequest("Invalid Signature"))
			return
		}

		if err := m.Datastore.QueryRowx(query).StructScan(&data); err != nil {
			json.NewEncoder(w).Encode(lib.ResponseInternalError(err))
			return
		}

		ok := lib.VerifySignature(token, signature, data.PrivateKey)
		if !ok {
			json.NewEncoder(w).Encode(lib.ResponseUnauthorized("Unauthorized!"))
			return
		}

		if lib.EncodeHMACSHA256(token, data.PrivateKey) != data.Signature {
			json.NewEncoder(w).Encode(lib.ResponseUnauthorized("Unauthorized!"))
			return
		}

		if data.Status < 1 {
			json.NewEncoder(w).Encode(lib.ResponseForbidden("Forbidden!"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
