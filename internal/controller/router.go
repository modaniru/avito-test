package controller

import (
	"encoding/json"
	log "log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/modaniru/avito/internal/service"
)

func NewRouter(r chi.Router, service *service.Service) {
	r.Mount("/user", NewUserRouter(service.User))
}

type ErrorResponse struct {
	Status       int    `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func writeError(w http.ResponseWriter, status int, err error) {
	b, err := json.Marshal(ErrorResponse{Status: status, ErrorMessage: err.Error()})
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
