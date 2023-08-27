package controller

import (
	"encoding/json"
	log "log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/modaniru/avito/internal/service"
)

type HistoryRouter struct {
	historyService service.History
}

func NewHistoryRouter(historyService service.History) chi.Router {
	h := HistoryRouter{historyService: historyService}
	r := chi.NewRouter()

	r.Get("/", h.GetHistory)
	r.Get("/date", h.GetHistoryByDate)

	return r
}

func (h *HistoryRouter) GetHistory(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	history, err := h.historyService.GetHistory(r.Context())
	if err != nil {
		log.Error("get history error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	b, err := json.Marshal(history)
	if err != nil {
		log.Error("marshal []history error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (h *HistoryRouter) GetHistoryByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	date := r.URL.Query().Get("date")
	//validate date

	history, err := h.historyService.GetHistoryByDate(r.Context(), date)
	if err != nil {
		log.Error("get history error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	b, err := json.Marshal(history)
	if err != nil {
		log.Error("marshal []history error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
