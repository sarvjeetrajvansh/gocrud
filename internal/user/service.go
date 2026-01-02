package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CREATE
func (s *Service) CreateUser(
	ctx context.Context,
	name, email string,
	age int,
) (User, error) {

	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.CreateUser")
	defer span.End()

	if name == "" || email == "" {
		err := errors.New("name and email are required")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return User{}, err
	}

	u := User{
		ID:    uuid.New(),
		Name:  name,
		Email: email,
		Age:   age,
	}

	user, err := s.repo.Create(ctx, u)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return user, err
}

// READ ALL
func (s *Service) GetUsers(ctx context.Context) ([]User, error) {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.GetUsers")
	defer span.End()

	users, err := s.repo.FindAll(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return users, err
}

// READ ONE
func (s *Service) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.GetUser")
	defer span.End()

	span.SetAttributes(attribute.String("user.id", id.String()))

	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return user, err
}

// UPDATE
func (s *Service) UpdateUser(
	ctx context.Context,
	id uuid.UUID,
	name, email string,
	age int,
) (User, error) {

	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.UpdateUser")
	defer span.End()

	span.SetAttributes(attribute.String("user.id", id.String()))

	if name == "" || email == "" {
		err := errors.New("name and email required")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return User{}, err
	}

	u := User{
		ID:    id,
		Name:  name,
		Email: email,
		Age:   age,
	}

	updated, err := s.repo.Update(ctx, u)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return updated, err
}

// DELETE
func (s *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	tracer := otel.Tracer("service.user")
	ctx, span := tracer.Start(ctx, "UserService.DeleteUser")
	defer span.End()

	span.SetAttributes(attribute.String("user.id", id.String()))

	err := s.repo.Delete(ctx, id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}
