package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type LogRecord struct {
	http.ResponseWriter
	status int
}

func (r *LogRecord) Write(p []byte) (int, error) {
	return r.ResponseWriter.Write(p)
}

func (r *LogRecord) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func WrapHandler(f http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		record := &LogRecord{
			ResponseWriter: w,
		}
		f.ServeHTTP(record, r)

		if record.status == http.StatusOK {
			log.Printf("200 OK %v", r)
		}

		if record.status == http.StatusBadRequest {
			log.Errorf("400 Bad Request %v", r)
		}

		if record.status == http.StatusConflict {
			log.Errorf("409 Conflict %v", r)
		}

		if record.status == http.StatusForbidden {
			log.Errorf("403 Forbidden %v", r)
		}

		if record.status == http.StatusNotFound {
			log.Errorf("404 Not Found! %v", r)
		}

		if record.status == http.StatusRequestTimeout {
			log.Errorf("Request Timeout %v", r)
		}

		if record.status == http.StatusUnauthorized {
			log.Errorf("401 Unauthorized %v", r)
		}

		if record.status == http.StatusInternalServerError {
			log.Errorf("500 Internal Server Error %v", r)
		}
	}
}
