package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/wardvisual/go-store-api/types"
	"github.com/wardvisual/go-store-api/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user types.RegisterUserPayload

	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
	}

	// TODO: Check if user exists
	// ...

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err = h.store.Create(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
