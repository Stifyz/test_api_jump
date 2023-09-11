package handlers

import (
	"encoding/json"

	"net/http"

	"test_api/internal/services"
)

type IUserHandler interface {
	GetAllUsers() http.HandlerFunc
}

type UserHandler struct {
	service services.IUserService
}

func NewUserHandler(service services.IUserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := h.service.GetAllUsers(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}
