package repositories

import (
	"github.com/ADMex1/GoProject/config"
	"github.com/ADMex1/GoProject/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindById(id uint) (*models.User, error)
	FindByPublicID(publicID string) (*models.User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user)
	return &user, err.Error
}

func (r *userRepository) FindById(id uint) (*models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByPublicID(publicID string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("public_id= ?", publicID).First(&user).Error
	return &user, err
}
