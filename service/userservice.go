package service

import (
	"context"
	"errors"
	"github.com/sarvjeetrajvansh/gocrud/models"
	"github.com/sarvjeetrajvansh/gocrud/storage"
	"go.opentelemetry.io/otel"
)

type Userservice struct {
	store *storage.Userstore
}

func NewUserservice(store *storage.Userstore) *Userservice {
	return &Userservice{store}
}

func (s *Userservice) CreateUser(ctx context.Context, name, email string, age int) (models.User, error) {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.CreateUser")
	defer span.End()

	if name == "" || email == "" {
		return models.User{}, errors.New("name and email are required")
	}
	user := models.User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	return s.store.Save(ctx, user), nil
}

func (s *Userservice) GetAllUsers(ctx context.Context) []models.User {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.GetAllUsers")
	defer span.End()

	return s.store.FindAll(ctx)
}

func (s *Userservice) GetUserByID(ctx context.Context, id string) (models.User, error) {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.GetUserByID")
	defer span.End()
	return s.store.FindByID(ctx, id)
}

func (s *Userservice) UpdateUser(ctx context.Context, id, name, email string, age int) (models.User, error) {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.UpdateUser")
	defer span.End()
	if name == "" || email == "" {
		return models.User{}, errors.New("name and email required")
	}
	return s.store.Update(ctx, id, models.User{
		Name:  name,
		Email: email,
		Age:   age,
	})
}

func (s *Userservice) DeleteUser(ctx context.Context, id string) error {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.DeleteUser")
	defer span.End()
	return s.store.Delete(ctx, id)
}
