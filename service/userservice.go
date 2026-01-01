package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/sarvjeetrajvansh/gocrud/models"
	"github.com/sarvjeetrajvansh/gocrud/storage"
	"go.opentelemetry.io/otel"
)

type Userservice struct {
	repo storage.UserRepository
}

func NewUserservice(repo storage.UserRepository) *Userservice {
	return &Userservice{repo: repo}
}

func (s *Userservice) CreateUser(ctx context.Context, name, email string, age int) (models.User, error) {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.CreateUser")
	defer span.End()

	if name == "" || email == "" {
		return models.User{}, errors.New("name and email are required")
	}
	user := models.User{
		ID:    uuid.New(),
		Name:  name,
		Email: email,
		Age:   age,
	}

	return s.repo.Create(ctx, user)
}

func (s *Userservice) GetAllUsers(ctx context.Context) ([]models.User, error) {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.GetAllUsers")
	defer span.End()

	return s.repo.FindAll(ctx)
}

func (s *Userservice) GetUserByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.GetUserByID")
	defer span.End()

	return s.repo.FindByID(ctx, id)
}

func (s *Userservice) UpdateUser(ctx context.Context, id uuid.UUID, name, email string, age int) (models.User, error) {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.UpdateUser")
	defer span.End()
	if name == "" || email == "" {
		return models.User{}, errors.New("name and email required")
	}
	return s.repo.Update(ctx, models.User{
		ID:    id,
		Name:  name,
		Email: email,
		Age:   age,
	})
}

func (s *Userservice) DeleteUser(ctx context.Context, id uuid.UUID) error {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.DeleteUser")
	defer span.End()
	return s.repo.Delete(ctx, id)
}
