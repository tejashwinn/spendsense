package repo

import (
	"spendsense/internal/models"

	"gorm.io/gorm"
)

type AccountRepo struct {
	DB *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *AccountRepo {
	return &AccountRepo{DB: db}
}

func (r *AccountRepo) CreateAccount(account *models.Account) error {
	return r.DB.Create(account).Error
}

func (r *AccountRepo) GetAccountByID(id interface{}) (*models.Account, error) {
	var account models.Account
	if err := r.DB.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepo) ListAccountsByUser(userID interface{}) ([]models.Account, error) {
	var accounts []models.Account
	if err := r.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *AccountRepo) UpdateAccount(account *models.Account) error {
	return r.DB.Save(account).Error
}

func (r *AccountRepo) DeleteAccount(id interface{}) error {
	return r.DB.Delete(&models.Account{}, id).Error
}

func (r *AccountRepo) GetAccountTypeByID(id interface{}) (*models.AccountType, error) {
	var accType models.AccountType
	if err := r.DB.First(&accType, id).Error; err != nil {
		return nil, err
	}
	return &accType, nil
}

func (r *AccountRepo) GetUserByID(id interface{}) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AccountRepo) GetCurrencyByID(id interface{}) (*models.Currency, error) {
	var currency models.Currency
	if err := r.DB.First(&currency, id).Error; err != nil {
		return nil, err
	}
	return &currency, nil
}
