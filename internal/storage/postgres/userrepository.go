package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sarvjeetrajvansh/gocrud/internal/user"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

// CREATE
func (r *UserRepo) Create(ctx context.Context, u user.User) (user.User, error) {
	err := r.db.WithContext(ctx).Create(&u).Error
	return u, err
}

// READ ALL
func (r *UserRepo) FindAll(ctx context.Context) ([]user.User, error) {
	var users []user.User
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}

// READ ONE
func (r *UserRepo) FindByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	var u user.User
	err := r.db.WithContext(ctx).First(&u, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user.User{}, errors.New("user not found")
	}
	return u, err
}

// UPDATE
func (r *UserRepo) Update(ctx context.Context, u user.User) (user.User, error) {
	result := r.db.WithContext(ctx).
		Model(&user.User{}).
		Where("id = ?", u.ID).
		Updates(map[string]interface{}{
			"name":  u.Name,
			"email": u.Email,
			"age":   u.Age,
		})

	if result.Error != nil {
		return user.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return user.User{}, errors.New("user not found")
	}

	return u, nil
}

// DELETE
func (r *UserRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Delete(&user.User{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

// CONSTRUCTOR
func NewUserRepo(db *gorm.DB) user.Repository {
	return &UserRepo{db: db}
}
