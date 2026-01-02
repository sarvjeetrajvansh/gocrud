package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, user User) (User, error)
	FindAll(ctx context.Context) ([]User, error)
	FindByID(ctx context.Context, id uuid.UUID) (User, error)
	Update(ctx context.Context, user User) (User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
