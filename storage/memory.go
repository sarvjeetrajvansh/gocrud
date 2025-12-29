package storage

import (
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
	s.nextID++
	return user
}
