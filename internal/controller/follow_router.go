package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (f *FollowRouter) FollowSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "read body error", err)
		return
	}

	var input FollowSegmentsInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unmarshal body error", err)
		return
	}

	err = validation.ValidateDate(*input.Expire)
	if err != nil {
		writeError(w, http.StatusBadRequest, "validate date error", err)
		return
	}

	err = f.userService.FollowToSegments(r.Context(), input.UserId, input.Segments, input.Expire)

	if err != nil {
		if errors.Is(err, repos.ErrUserAlreadyHasThisSegment) {
			writeError(w, http.StatusBadRequest, "user alredy has some segments in this list", err)
			return
		} else if errors.Is(err, repos.ErrSegmentNotFound) {
			writeError(w, http.StatusBadRequest, "some segments in list not exist", err)
			return
		}
		writeError(w, http.StatusInternalServerError, "follow to segments error", err)
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
		writeError(w, http.StatusBadRequest, "read body error", err)
		return
	}

	var input UnfollowSegmentsInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unmarshal body error", err)
		return
	}

	err = f.userService.UnFollowToSegments(r.Context(), input.UserId, input.Segments)
	if err != nil {
		if errors.Is(err, repos.ErrUserOrSegmentNotExists) {
			writeError(w, http.StatusBadRequest, "user or segment in list are not exist", err)
			return
		}
		writeError(w, http.StatusInternalServerError, "unfollow to segments error", err)
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
		writeError(w, http.StatusBadRequest, "read body error", err)
		return
	}

	var input GetUserSegmentsInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unmarshal body error", err)
		return
	}

	segments, err := f.userService.GetUserSegments(r.Context(), input.Id)
	if err != nil {
		if errors.Is(err, repos.ErrUserNotFound) {
			writeError(w, http.StatusNotFound, fmt.Sprintf("user with id=%d was not found", input.Id), err)
			return
		}
		writeError(w, http.StatusBadRequest, "get user segments error", err)
		return
	}
	b, err = json.Marshal(segments)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "marshal []segments error", err)
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
		writeError(w, http.StatusBadRequest, "read body error", err)
		return
	}

	var input RandomFollowInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unmarshal body error", err)
		return
	}
	err = validation.ValidatePercent(input.Percent)
	if err != nil {
		writeError(w, http.StatusBadRequest, "validate percent error", err)
		return
	}

	rowsAffected, err := f.userService.FollowRandomUsers(r.Context(), input.Name, input.Percent)

	if err != nil {
		if errors.Is(err, repos.ErrUserAlreadyHasThisSegment) {
			writeError(w, http.StatusBadRequest, "user alredy has some segments in this list", err)
			return
		} else if errors.Is(err, repos.ErrUserOrSegmentNotExists) {
			writeError(w, http.StatusBadRequest, "user or some segments in list not exist", err)
			return
		}
		writeError(w, http.StatusInternalServerError, "follow to segments error", err)
		return
	}

	type RandomFollowResponse struct {
		RowsAffected int `json:"rows_affected"`
	}

	b, err = json.Marshal(&RandomFollowResponse{RowsAffected: rowsAffected})
	if err != nil {
		writeError(w, http.StatusInternalServerError, "marshal RandomFollowResponse error", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
