package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	log "log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/modaniru/avito/internal/service"
	"github.com/modaniru/avito/internal/storage/repos"
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

func (s *SegmentRouter) SaveSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("read body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var input SaveSegmentInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		log.Error("unmarshal body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	id, err := s.segmentService.SaveSegment(r.Context(), input.Name)
	if err != nil {
		log.Error("save user error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	type response struct {
		Id int `json:"id"`
	}
	b, err = json.Marshal(response{Id: id})
	if err != nil {
		log.Error("marshal users error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

type DeleteSegmentInput struct {
	Name string `json:"name"`
}

func (s *SegmentRouter) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("read body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var input DeleteSegmentInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		log.Error("unmarshal body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	err = s.segmentService.DeleteSegment(r.Context(), input.Name)
	if err != nil {
		if errors.Is(err, repos.ErrSegmentNotFound) {
			log.Error("user not found error", log.String("error", err.Error()))
			writeError(w, http.StatusNotFound, errors.New(fmt.Sprintf("segment with name=%s was not found", input.Name)))
			return
		}
		log.Error("delete segment error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *SegmentRouter) GetSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	segments, err := s.segmentService.GetSegments(r.Context())
	if err != nil {
		log.Error("get segments error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}
	b, err := json.Marshal(segments)
	if err != nil {
		log.Error("marshal segments error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
