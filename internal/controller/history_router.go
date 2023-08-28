package controller

import (
	"encoding/json"
	"errors"
	log "log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/modaniru/avito/internal/service"
	"github.com/modaniru/avito/internal/service/services"
	"github.com/modaniru/avito/internal/validation"
)

type HistoryRouter struct {
	historyService service.History
}

func NewHistoryRouter(historyService service.History) chi.Router {
	h := HistoryRouter{historyService: historyService}
	r := chi.NewRouter()

	r.Get("/date", h.GetHistoryByDate)

	return r
}

func (h *HistoryRouter) GetHistoryByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	date := r.URL.Query().Get("date")
	err := validation.ValidateDate(date)
	if err != nil {
		log.Error("validate date error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, errors.New("validate date error"))
		return
	}

	history, err := h.historyService.GetHistoryByDate(r.Context(), date)
	if err != nil {
		if errors.Is(err, services.ErrServiceUnavailible) {
			log.Error("history service is unavailible", log.String("error", err.Error()))
			writeError(w, http.StatusBadRequest, errors.New("history service is unavailible"))
			return
		}
		log.Error("get history error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	type historyResponse struct {
		Link string `json:"link"`
	}

	b, err := json.Marshal(&historyResponse{Link: history})
	if err != nil {
		log.Error("marshal history error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
