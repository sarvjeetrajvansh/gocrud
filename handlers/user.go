package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/sarvjeetrajvansh/gocrud/service"
	"net/http"
	"strconv"
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
			ID:    strconv.Itoa(int(user.ID)),
			Name:  user.Name,
			Email: user.Email,
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
		id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)

		user, err := userService.GetUserByID(r.Context(), uint(id))
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

		id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
		if id < 0 {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		var req CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		updated, err := userService.UpdateUser(r.Context(), uint(id), req.Name, req.Email, req.Age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := UserResponse{
			ID:    strconv.Itoa(int(updated.ID)),
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

		id, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
		if id < 0 {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		if err := userService.DeleteUser(r.Context(), uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
