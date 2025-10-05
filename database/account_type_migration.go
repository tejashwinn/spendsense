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
	defaultTypes := map[string]string{
		"BANK":         "Bank",
		"CASH":         "Cash",
		"WALLET":       "Wallet",
		"CREDIT_CARD":  "Credit Card",
		"LOAN":         "Loan",
		"INVESTMENT":   "Investment",
		"RETIREMENT":   "Retirement",
		"CRYPTO":       "Crypto",
		"LIABILITY":    "Liablity",
		"REIMBURSABLE": "Reimbursable",
	}

	for k, v := range defaultTypes {
		var count int64
		db.Model(&models.AccountType{}).Where("type = ?", k).Count(&count)
		if count == 0 {
			db.Create(&models.AccountType{Type: k, Name: v})
		}
	}

	return nil
}
