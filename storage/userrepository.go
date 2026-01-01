package storage

import (
	"context"
	"github.com/google/uuid"
	"github.com/sarvjeetrajvansh/gocrud/models"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (models.User, error)
	FindAll(ctx context.Context) ([]models.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (models.User, error)
	Update(ctx context.Context, user models.User) (models.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
