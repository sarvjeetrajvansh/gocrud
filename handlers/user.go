package handlers

import (
	"encoding/json"
	"github.com/sarvjeetrajvansh/gocrud/service"
	"net/http"
	"strings"
)

func getIDFromPath(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return ""
	}
	return parts[2]
}

func Users(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			users := userService.GetAllUsers()
			json.NewEncoder(w).Encode(users)

		case http.MethodPost:
			var req CreateUserRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "invalid body", http.StatusBadRequest)
				return
			}

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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(resp)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func UserByID(userService *service.Userservice) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := getIDFromPath(r.URL.Path)
		if id == "" {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		switch r.Method {

		case http.MethodGet:
			user, err := userService.GetUserByID(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(user)

		case http.MethodPut:
			var req CreateUserRequest
			json.NewDecoder(r.Body).Decode(&req)

			user, err := userService.UpdateUser(id, req.Name, req.Email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(user)

		case http.MethodDelete:
			if err := userService.DeleteUser(id); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
