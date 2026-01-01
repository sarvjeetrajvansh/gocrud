package storage

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/sarvjeetrajvansh/gocrud/models"
	"go.opentelemetry.io/otel"
)

type Userstore struct {
	users map[uuid.UUID]models.User
}

func NewUserstore() *Userstore {
	return &Userstore{
		users: make(map[uuid.UUID]models.User),
	}
}

func (s *Userstore) Save(ctx context.Context, user models.User) models.User {
	tracer := otel.Tracer("storage.user")
	_, span := tracer.Start(ctx, "UserStore.Save")
	defer span.End()

	user.ID = uuid.New()
	s.users[user.ID] = user
	return user
}

func (s *Userstore) FindAll(ctx context.Context) []models.User {
	tracer := otel.Tracer("storage.user")
	_, span := tracer.Start(ctx, "UserStore.FindAll")
	defer span.End()

	users := make([]models.User, 0)
	for _, u := range s.users {
		users = append(users, u)
	}
	return users
}

func (s *Userstore) FindByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	tracer := otel.Tracer("storage.user")
	_, span := tracer.Start(ctx, "UserStore.FindByID")
	defer span.End()
	user, ok := s.users[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *Userstore) Update(ctx context.Context, id uuid.UUID, updated models.User) (models.User, error) {
	tracer := otel.Tracer("storage.user")
	_, span := tracer.Start(ctx, "UserStore.Update")
	defer span.End()

	_, ok := s.users[id]
	if !ok {
		return models.User{}, errors.New("user not found")
	}
	
	return updated, nil
}

func (s *Userstore) Delete(ctx context.Context, id uuid.UUID) error {
	tracer := otel.Tracer("storage.user")
	_, span := tracer.Start(ctx, "UserStore.Delete")
	defer span.End()

	if _, ok := s.users[id]; !ok {
		return errors.New("user not found")
	}
	delete(s.users, id)
	return nil
}
