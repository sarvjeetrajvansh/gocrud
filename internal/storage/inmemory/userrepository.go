package inmemory

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sarvjeetrajvansh/gocrud/internal/user"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type UserRepo struct {
	users map[uuid.UUID]user.User
}

func NewUserRepo() user.Repository {
	return &UserRepo{
		users: make(map[uuid.UUID]user.User),
	}
}

// CREATE
func (r *UserRepo) Create(ctx context.Context, u user.User) (user.User, error) {
	tracer := otel.Tracer("storage.memory")
	_, span := tracer.Start(ctx, "UserRepo.Create")
	defer span.End()

	if _, exists := r.users[u.ID]; exists {
		err := errors.New("user already exists")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return user.User{}, err
	}

	r.users[u.ID] = u
	return u, nil
}

// READ ALL
func (r *UserRepo) FindAll(ctx context.Context) ([]user.User, error) {
	tracer := otel.Tracer("storage.memory")
	_, span := tracer.Start(ctx, "UserRepo.FindAll")
	defer span.End()

	users := make([]user.User, 0, len(r.users))
	for _, u := range r.users {
		users = append(users, u)
	}

	return users, nil
}

// READ ONE
func (r *UserRepo) FindByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	tracer := otel.Tracer("storage.memory")
	_, span := tracer.Start(ctx, "UserRepo.FindByID")
	defer span.End()

	span.SetAttributes(attribute.String("user.id", id.String()))

	u, ok := r.users[id]
	if !ok {
		err := errors.New("user not found")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return user.User{}, err
	}

	return u, nil
}

// UPDATE
func (r *UserRepo) Update(ctx context.Context, u user.User) (user.User, error) {
	tracer := otel.Tracer("storage.memory")
	_, span := tracer.Start(ctx, "UserRepo.Update")
	defer span.End()

	span.SetAttributes(attribute.String("user.id", u.ID.String()))

	if _, ok := r.users[u.ID]; !ok {
		err := errors.New("user not found")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return user.User{}, err
	}

	r.users[u.ID] = u
	return u, nil
}

// DELETE
func (r *UserRepo) Delete(ctx context.Context, id uuid.UUID) error {
	tracer := otel.Tracer("storage.memory")
	_, span := tracer.Start(ctx, "UserRepo.Delete")
	defer span.End()

	span.SetAttributes(attribute.String("user.id", id.String()))

	if _, ok := r.users[id]; !ok {
		err := errors.New("user not found")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	delete(r.users, id)
	return nil
}
