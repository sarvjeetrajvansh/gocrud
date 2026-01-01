package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sarvjeetrajvansh/gocrud/service"
	"net/http"
)

func GetUsers(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, _ := userService.GetAllUsers(r.Context())
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			return
		}
	}
}
func CreateUser(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			return
		}

		user, err := userService.CreateUser(r.Context(), req.Name, req.Email, req.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := UserResponse{
			ID:    uuid.New().String(),
			Name:  user.Name,
			Email: user.Email,
			Age:   user.Age,
		}

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			return
		}
	}
}
func GetUserByID(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))

		user, err := userService.GetUserByID(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			return
		}
	}
}
func UpdateUser(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "invalid uuid", http.StatusBadRequest)
			return
		}

		var req CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		updated, err := userService.UpdateUser(r.Context(), id, req.Name, req.Email, req.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := UserResponse{
			Name:  updated.Name,
			Email: updated.Email,
			Age:   updated.Age,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			return
		}
	}
}
func DeleteUser(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "invalid uuid", http.StatusBadRequest)
			return
		}

		if err := userService.DeleteUser(r.Context(), id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
