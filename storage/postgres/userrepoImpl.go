package postgres

import (
	"context"
	"errors"
	"github.com/sarvjeetrajvansh/gocrud/models"
	"github.com/sarvjeetrajvansh/gocrud/storage"
	"gorm.io/gorm"
)

type userRepoImpl struct {
	db *gorm.DB
}

func (r *userRepoImpl) Create(ctx context.Context, user models.User) (models.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	return user, err
}

func (r *userRepoImpl) FindAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}

func (r *userRepoImpl) FindByID(ctx context.Context, id uint) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user not found")
	}
	return user, err
}

func (r *userRepoImpl) Update(ctx context.Context, user models.User) (models.User, error) {
	err := r.db.WithContext(ctx).Save(&user).Error
	return user, err
}

func (r *userRepoImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}

func NewuserRepo(db *gorm.DB) storage.UserRepository {
	return &userRepoImpl{db: db}
}
