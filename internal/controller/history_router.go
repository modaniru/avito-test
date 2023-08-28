package controller

import (
	"encoding/json"
	"errors"
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

	r.Get("/", h.GetHistoryByDate)

	return r
}

type GetHistoryResponse struct {
	Link string `json:"link"`
}

// @Summary		get history
// @Tags			history
// @Description	Получить историю добавления/удаления сегментов у пользователя (ГГГГ-ММ-ДД или ГГГГ-ММ)
// @Accept			json
// @Produce		json
// @Param			date	query		string	true	"get history by date YYYY-MM or YYYY-MM-DD"
// @Success		200		{object}	GetHistoryResponse
// @Falure			400 {object} ErrorResponse 1
// @Falure			500 {object} ErrorResponse 1
// @Router			/history/ [get]
func (h *HistoryRouter) GetHistoryByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	date := r.URL.Query().Get("date")
	err := validation.ValidateDate(date)
	if err != nil {
		writeError(w, http.StatusBadRequest, "validate date error", err)
		return
	}

	history, err := h.historyService.GetHistoryByDate(r.Context(), date)
	if err != nil {
		if errors.Is(err, services.ErrServiceUnavailible) {
			writeError(w, http.StatusBadRequest, "history service is unavailible", err)
			return
		}
		writeError(w, http.StatusInternalServerError, "get history error", err)
		return
	}

	b, err := json.Marshal(&GetHistoryResponse{Link: history})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "marshal history error", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
