package storage

import (
	"errors"
	"github.com/sarvjeetrajvansh/gocrud/models"
	"strconv"
)

type Userstore struct {
	users  map[string]models.User
	nextID int
}

func NewUserstore() *Userstore {
	return &Userstore{
		users:  make(map[string]models.User),
		nextID: 1,
	}
}

func (s *Userstore) Save(user models.User) models.User {
	user.ID = strconv.Itoa(s.nextID)
	s.users[user.ID] = user
	s.nextID++
	return user
}

func (s *Userstore) FindAll() []models.User {
	users := make([]models.User, 0)
	for _, u := range s.users {
		users = append(users, u)
	}
	return users
}

func (s *Userstore) FindByID(id string) (models.User, error) {
	user, ok := s.users[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *Userstore) Update(id string, updated models.User) (models.User, error) {
	_, ok := s.users[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}
	updated.ID = id
	s.users[id] = updated
	return updated, nil
}

func (s *Userstore) Delete(id string) error {
	if _, ok := s.users[id]; !ok {
		return errors.New("user not found")
	}
	delete(s.users, id)
	return nil
}
