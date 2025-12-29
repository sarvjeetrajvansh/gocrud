package service

import (
	"errors"
	"github.com/sarvjeetrajvansh/gocrud/models"
	"github.com/sarvjeetrajvansh/gocrud/storage"
)

type Userservice struct {
	store *storage.Userstore
}

func NewUserservice(store *storage.Userstore) *Userservice {
	return &Userservice{store}
}

func (s *Userservice) CreateUser(name, email string, age int) (models.User, error) {
	if name == "" || email == "" {
		return models.User{}, errors.New("name and email are required")
	}
	user := models.User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	return s.store.Save(user), nil
}
