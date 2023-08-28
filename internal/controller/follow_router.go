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
	"github.com/modaniru/avito/internal/validation"
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
	r.Post("/auto", f.RandomFollow)

	return r
}

type FollowSegmentsInput struct {
	UserId   int      `json:"user_id"`
	Segments []string `json:"segments"`
	Expire   *string  `json:"expire,omitempty"`
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

	err = validation.ValidateDate(*input.Expire)
	if err != nil {
		log.Error("validate date error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, errors.New("validate date error"))
		return
	}

	err = f.userService.FollowToSegments(r.Context(), input.UserId, input.Segments, input.Expire)

	if err != nil {
		if errors.Is(err, repos.ErrUserAlreadyHasThisSegment) {
			log.Error("user alredy has some segments in this list", log.String("error", err.Error()))
			writeError(w, http.StatusBadRequest, errors.New("user alredy has some segments in this list"))
			return
		} else if errors.Is(err, repos.ErrUserOrSegmentNotExists) {
			log.Error("user or some segments in list not exist", log.String("error", err.Error()))
			writeError(w, http.StatusBadRequest, errors.New("user or some segments in list not exist"))
			return
		}
		log.Error("follow to segments error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
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
		if errors.Is(err, repos.ErrUserOrSegmentNotExists) {
			log.Error("user or segment in list are not exist", log.String("error", err.Error()))
			writeError(w, http.StatusBadRequest, errors.New("user or segment in list are not exist"))
			return
		}
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
		if errors.Is(err, repos.ErrUserNotFound) {
			log.Error("user not found error", log.String("error", err.Error()))
			writeError(w, http.StatusNotFound, fmt.Errorf("user with id=%d was not found", input.Id))
			return
		}
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

type RandomFollowInput struct {
	Name    string  `json:"name"`
	Percent float64 `json:"percent"`
}

func (f *FollowRouter) RandomFollow(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("read body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}

	var input RandomFollowInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		log.Error("unmarshal body error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, err)
		return
	}
	err = validation.ValidatePercent(input.Percent)
	if err != nil {
		log.Error("validate percent error", log.String("error", err.Error()))
		writeError(w, http.StatusBadRequest, errors.New("validate percent error"))
		return
	}

	rowsAffected, err := f.userService.FollowRandomUsers(r.Context(), input.Name, input.Percent)

	if err != nil {
		if errors.Is(err, repos.ErrUserAlreadyHasThisSegment) {
			log.Error("user alredy has some segments in this list", log.String("error", err.Error()))
			writeError(w, http.StatusBadRequest, errors.New("user alredy has some segments in this list"))
			return
		} else if errors.Is(err, repos.ErrUserOrSegmentNotExists) {
			log.Error("user or some segments in list not exist", log.String("error", err.Error()))
			writeError(w, http.StatusBadRequest, errors.New("user or some segments in list not exist"))
			return
		}
		log.Error("follow to segments error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	type RandomFollowResponse struct {
		RowsAffected int `json:"rows_affected"`
	}

	b, err = json.Marshal(&RandomFollowResponse{RowsAffected: rowsAffected})
	if err != nil {
		log.Error("marshal RandomFollowResponse error", log.String("error", err.Error()))
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
