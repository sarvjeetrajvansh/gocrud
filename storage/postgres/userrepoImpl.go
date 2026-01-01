package postgres

import (
	"context"
	"errors"
	"github.com/google/uuid"
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

func (r *userRepoImpl) FindByID(ctx context.Context, id uuid.UUID) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user not found")
	}
	return user, err
}

func (r *userRepoImpl) Update(
	ctx context.Context,
	user models.User,
) (models.User, error) {

	result := r.db.WithContext(ctx).
		Model(&models.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
			"age":   user.Age,
		})

	if result.Error != nil {
		return models.User{}, result.Error
	}

	if result.RowsAffected == 0 {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *userRepoImpl) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).
		Delete(&models.User{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}

func NewuserRepo(db *gorm.DB) storage.UserRepository {
	return &userRepoImpl{db: db}
}
