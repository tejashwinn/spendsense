package database

import (
	"github.com/tejashwinn/spendsense/internal/models"
	"gorm.io/gorm"
)

func MigrateCurrency(db *gorm.DB) error {
	// Auto migrate the table
	if err := db.AutoMigrate(&models.Currency{}); err != nil {
		return err
	}

	defaultCurrencies := []struct {
		Code     string
		Name     string
		Symbol   string
		Decimals int
	}{
		{"USD", "United States Dollar", "$", 2},
		{"EUR", "Euro", "€", 2},
		{"GBP", "British Pound Sterling", "£", 2},
		{"JPY", "Japanese Yen", "¥", 0},
		{"INR", "Indian Rupee", "₹", 2},
		{"AUD", "Australian Dollar", "A$", 2},
		{"CAD", "Canadian Dollar", "C$", 2},
		{"CHF", "Swiss Franc", "CHF", 2},
		{"CNY", "Chinese Yuan", "¥", 2},
		{"SGD", "Singapore Dollar", "S$", 2},
		{"HKD", "Hong Kong Dollar", "HK$", 2},
		{"NZD", "New Zealand Dollar", "NZ$", 2},
		{"SEK", "Swedish Krona", "kr", 2},
		{"NOK", "Norwegian Krone", "kr", 2},
		{"DKK", "Danish Krone", "kr", 2},
		{"ZAR", "South African Rand", "R", 2},
		{"BRL", "Brazilian Real", "R$", 2},
		{"RUB", "Russian Ruble", "₽", 2},
		{"MXN", "Mexican Peso", "$", 2},
		{"AED", "UAE Dirham", "د.إ", 2},
	}

	for _, c := range defaultCurrencies {
		var count int64
		db.Model(&models.Currency{}).Where("code = ?", c.Code).Count(&count)
		if count == 0 {
			db.Create(&models.Currency{
				Code:          c.Code,
				Name:          c.Name,
				Symbol:        c.Symbol,
				DecimalPlaces: c.Decimals,
			})
		}
	}

	return nil
}
