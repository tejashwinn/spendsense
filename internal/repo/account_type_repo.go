package repo

import (
	"github.com/tejashwinn/spendsense/internal/models"

	"gorm.io/gorm"
)

type AccountTypeRepo struct {
	DB *gorm.DB
}

func NewAccountTypeRepo(db *gorm.DB) *AccountTypeRepo {
	return &AccountTypeRepo{DB: db}
}

func (r *AccountTypeRepo) ListAccountTypes() ([]models.AccountType, error) {
	var accountTypes []models.AccountType
	if err := r.DB.Find(&accountTypes).Error; err != nil {
		return nil, err
	}
	return accountTypes, nil
}

func (r *AccountTypeRepo) GetAccountTypeByID(id interface{}) (*models.AccountType, error) {
	var accountType models.AccountType
	if err := r.DB.First(&accountType, id).Error; err != nil {
		return nil, err
	}
	return &accountType, nil
}
