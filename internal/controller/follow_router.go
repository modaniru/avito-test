package controller

import (
	"encoding/json"
	"io"
	log "log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/modaniru/avito/internal/service"
)

type FollowRouter struct {
	userService service.User
}

func NewFollowRouter(userService service.User) chi.Router {
	f := FollowRouter{userService: userService}

	r := chi.NewRouter()
	r.Post("/", f.FollowSegments)
	r.Delete("/", f.UnfollowSegments)
	r.Get("/", f.GetUserSegments)

	return r
}

type FollowSegmentsInput struct {
	UserId   int      `json:"user_id"`
	Segments []string `json:"segments"`
}

// TODO user not found or segment handle error
func (f *FollowRouter) FollowSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("read body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var input FollowSegmentsInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		log.Error("unmarshal body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	err = f.userService.FollowToSegments(r.Context(), input.UserId, input.Segments)

	if err != nil {
		log.Error("follow to segments error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type UnfollowSegmentsInput struct {
	UserId   int      `json:"user_id"`
	Segments []string `json:"segments"`
}

// TODO handle user or segment not found error
func (f *FollowRouter) UnfollowSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("read body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var input UnfollowSegmentsInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		log.Error("unmarshal body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	err = f.userService.UnFollowToSegments(r.Context(), input.UserId, input.Segments)
	if err != nil {
		log.Error("unfollow to segments error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type GetUserSegmentsInput struct {
	Id int `json:"id"`
}

func (f *FollowRouter) GetUserSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("read body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var input GetUserSegmentsInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		log.Error("unmarshal body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	segments, err := f.userService.GetUserSegments(r.Context(), input.Id)
	if err != nil {
		log.Error("get user segments error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}
	b, err = json.Marshal(segments)
	if err != nil {
		log.Error("marshal []segments error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
