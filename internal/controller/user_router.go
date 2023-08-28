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
)

type UserRouter struct {
	userService service.User
}

func NewUserRouter(userService service.User) chi.Router {
	u := UserRouter{userService: userService}
	r := chi.NewRouter()

	r.Post("/", u.SaveUser)
	r.Delete("/", u.DeleteUser)
	r.Get("/all", u.GetUsers)
	r.Mount("/segment", NewFollowRouter(userService))
	return r
}

type SaveUserInput struct {
	Id int `json:"id"`
}

// @Summary		save user id
// @Tags			user
// @Description	Сохранить пользователя
// @Accept			json
// @Produce		json
// @Param			input	body	SaveUserInput	true	"user id"
// @Success		201
// @Falure			400 {object} ErrorResponse 1
// @Falure			500 {object} ErrorResponse 1
// @Router			/user/ [post]
func (u *UserRouter) SaveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "read body error", err)
		return
	}

	var input SaveUserInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unmarshal body error", err)
		return
	}

	err = u.userService.SaveUser(r.Context(), input.Id)
	if err != nil {
		if errors.Is(err, repos.ErrUserAlreadyExists) {
			writeError(w, http.StatusBadRequest, fmt.Sprintf("user with id=%d already exists", input.Id), err)
			return
		}
		writeError(w, http.StatusInternalServerError, "save user error", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type DeleteUserInput struct {
	Id int `json:"id"`
}

// @Summary		delete user by id
// @Tags			user
// @Description	Удалить пользователя
// @Accept			json
// @Produce		json
// @Param			input	body	DeleteUserInput	true	"user id"
// @Success		204
// @Falure			400 {object} ErrorResponse 1
// @Falure			404 {object} ErrorResponse 1
// @Falure			500 {object} ErrorResponse 1
// @Router			/user/ [delete]
func (u *UserRouter) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		writeError(w, http.StatusBadRequest, "read body error", err)
		return
	}

	var input DeleteUserInput
	err = json.Unmarshal(b, &input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "unmarshal body error", err)
		return
	}

	err = u.userService.DeleteUser(r.Context(), input.Id)
	if err != nil {
		if errors.Is(err, repos.ErrUserNotFound) {
			writeError(w, http.StatusNotFound, fmt.Sprintf("user with id=%d was not found", input.Id), err)
			return
		}
		writeError(w, http.StatusInternalServerError, "delete user error", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary		get users
// @Tags			user
// @Description	Получить всех пользователей
// @Accept			json
// @Produce		json
// @Success		200	{array}	entity.User
// @Falure			400 {object} ErrorResponse 1
// @Falure			500 {object} ErrorResponse 1
// @Router			/user/all [get]
func (u *UserRouter) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	users, err := u.userService.GetUsers(r.Context())
	if err != nil {
		writeError(w, http.StatusBadRequest, "get users error", err)
		return
	}
	b, err := json.Marshal(users)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "marshal users error", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
