package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/modaniru/avito/internal/service"
	"github.com/modaniru/avito/internal/storage/repos"
	"github.com/modaniru/avito/internal/validation"
)

type SegmentRouter struct {
	segmentService service.Segment
}

func NewSegmentRouter(segmentService service.Segment) chi.Router {
	s := SegmentRouter{segmentService: segmentService}

	r := chi.NewRouter()
	r.Post("/", s.SaveSegment)
	r.Get("/all", s.GetSegments)
	r.Delete("/", s.DeleteSegment)

	return r
}

type SaveSegmentInput struct {
	Name string `json:"name"`
}

type SaveSegmentResponse struct {
	Id int `json:"id"`
}

// @Summary		save segment
// @Tags			segment
// @Description	Сохранить сегмент
// @Accept			json
// @Produce		json
// @Param			input	body		SaveSegmentInput	true	"segment name"
// @Success		201		{object}	SaveSegmentResponse
// @Falure			400 {object} ErrorResponse 1
// @Falure			500 {object} ErrorResponse 1
// @Router			/segment/ [post]
func (s *SegmentRouter) SaveSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "read body error", err)
		return
	}

	var input SaveSegmentInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unmarshal body error", err)
		return
	}
	err = validation.ValidateSegment(input.Name)
	if err != nil {
		writeError(w, http.StatusBadRequest, "validate segment name error", err)
		return
	}

	id, err := s.segmentService.SaveSegment(r.Context(), input.Name)
	if err != nil {
		if errors.Is(err, repos.ErrSegmentAlreadyExists) {
			writeError(w, http.StatusBadRequest, "segment already exists error", err)
			return
		}
		writeError(w, http.StatusBadRequest, "save user error", err)
		return
	}

	b, err = json.Marshal(SaveSegmentResponse{Id: id})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "marshal users error", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

type DeleteSegmentInput struct {
	Name string `json:"name"`
}

// @Summary		delete segment by name
// @Tags			segment
// @Description	Удалить сегмент по его имени
// @Accept			json
// @Produce		json
// @Param			input	body	DeleteSegmentInput	true	"segment name"
// @Success		204
// @Falure			400 {object} ErrorResponse 1
// @Falure			404 {object} ErrorResponse 1
// @Falure			500 {object} ErrorResponse 1
// @Router			/segment/ [delete]
func (s *SegmentRouter) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "read body error", err)
		return
	}

	var input DeleteSegmentInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unmarshal body error", err)
		return
	}

	err = s.segmentService.DeleteSegment(r.Context(), input.Name)
	if err != nil {
		if errors.Is(err, repos.ErrSegmentNotFound) {
			writeError(w, http.StatusNotFound, "segment not found error", err)
			return
		}
		writeError(w, http.StatusBadRequest, "delete segment error", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary		get segments
// @Tags			segment
// @Description	Получить все сегменты
// @Accept			json
// @Produce		json
// @Success		200	{array}	entity.Segment
// @Falure			400 {object} ErrorResponse 1
// @Falure			500 {object} ErrorResponse 1
// @Router			/segment/all [get]
func (s *SegmentRouter) GetSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	segments, err := s.segmentService.GetSegments(r.Context())
	if err != nil {
		writeError(w, http.StatusBadRequest, "get segments error", err)
		return
	}
	b, err := json.Marshal(segments)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "marshal segments error", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
