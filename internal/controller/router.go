package controller

import (
	"encoding/json"
	log "log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/modaniru/avito/internal/service"
)

func NewRouter(r chi.Router, service *service.Service) {
	r.Use(LoggingMiddleware)
	r.Mount("/user", NewUserRouter(service.User))
	r.Mount("/segment", NewSegmentRouter(service.Segment))
	r.Mount("/history", NewHistoryRouter(service.History))
}

type ErrorResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func writeError(w http.ResponseWriter, status int, message string, err error) {
	log.Error(message, log.String("error", err.Error()))
	b, err := json.Marshal(ErrorResponse{Status: status, ErrorMessage: message})
	if err != nil {
		log.Error("marshar 'ErrorResponse' error", log.String("error", err.Error()))
		return
	}
	w.WriteHeader(status)
	_, err = w.Write(b)
	if err != nil {
		log.Error("write to ResponseWriter error", log.String("error", err.Error()))
	}
}
