package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/sarvjeetrajvansh/gocrud/service"
	"net/http"
)

func GetUsers(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := userService.GetAllUsers()
		json.NewEncoder(w).Encode(users)
	}
}
func CreateUser(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateUserRequest
		json.NewDecoder(r.Body).Decode(&req)

		user, err := userService.CreateUser(req.Name, req.Email, req.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}
func GetUserByID(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		user, err := userService.GetUserByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}
func UpdateUser(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		var req CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		updated, err := userService.UpdateUser(id, req.Name, req.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := UserResponse{
			ID:    updated.ID,
			Name:  updated.Name,
			Email: updated.Email,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
func DeleteUser(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		if err := userService.DeleteUser(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
