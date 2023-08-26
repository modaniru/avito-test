package controller

import (
	"net/http"

	"github.com/urfave/negroni"
	log "log/slog"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		lrw := negroni.NewResponseWriter(w)
		next.ServeHTTP(lrw, r)

		log.Info("request", log.String("method", r.Method), log.String("uri", r.RequestURI), log.Int("status", lrw.Status()))
	})
}
