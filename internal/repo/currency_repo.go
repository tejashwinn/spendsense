package repo

import (
	"spendsense/internal/models"

	"gorm.io/gorm"
)

type CurrencyRepo struct {
	DB *gorm.DB
}

func NewCurrencyRepo(db *gorm.DB) *CurrencyRepo {
	return &CurrencyRepo{DB: db}
}

func (r *CurrencyRepo) ListCurrencies() ([]models.Currency, error) {
	var currencies []models.Currency
	if err := r.DB.Find(&currencies).Error; err != nil {
		return nil, err
	}
	return currencies, nil
}

func (r *CurrencyRepo) GetCurrencyByID(id interface{}) (*models.Currency, error) {
	var currency models.Currency
	if err := r.DB.First(&currency, id).Error; err != nil {
		return nil, err
	}
	return &currency, nil
}
