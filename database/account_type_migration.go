package database

import (
	"github.com/tejashwinn/spendsense/internal/models"
	"gorm.io/gorm"
)

func MigrateAccountType(db *gorm.DB) error {
	// Auto migrate the table
	if err := db.AutoMigrate(&models.AccountType{}); err != nil {
		return err
	}

	// Seed master data if not exists
	defaultTypes := []string{
		"BANK",
		"CASH",
		"WALLET",
		"CREDIT_CARD",
		"LOAN",
		"INVESTMENT",
		"RETIREMENT",
		"CRYPTO",
		"LIABILITY",
		"REIMBURSABLE",
	}

	for _, t := range defaultTypes {
		var count int64
		db.Model(&models.AccountType{}).Where("name = ?", t).Count(&count)
		if count == 0 {
			db.Create(&models.AccountType{Name: t})
		}
	}

	return nil
}
