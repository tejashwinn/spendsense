package repo

import (
	"github.com/tejashwinn/spendsense/internal/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) CreateUser(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepo) GetUserByID(id interface{}) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) UpdateUser(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepo) DeleteUser(id interface{}) error {
	return r.DB.Delete(&models.User{}, id).Error
}
