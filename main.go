package main

import (
	"net/http"

	"github.com/sarvjeetrajvansh/gocrud/handlers"
	"github.com/sarvjeetrajvansh/gocrud/service"
	"github.com/sarvjeetrajvansh/gocrud/storage"
)

const PORT string = ":8080"

func main() {
	userStore := storage.NewUserstore()
	userService := service.NewUserservice(userStore)

	http.HandleFunc("/users", handlers.CreateUser(userService))

	http.HandleFunc("/", handlers.HelloUser())
	http.ListenAndServe(PORT, nil)

}
