package database

import (
	"log"

	"github.com/tejashwinn/spendsense/config"
	"github.com/tejashwinn/spendsense/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormDatabase is a wrapper for the gorm framework.
type GormDatabase struct {
	DB *gorm.DB
}

func New(cfg *config.Config) (*GormDatabase, error) {
	if cfg.DBUrl == "" {
		log.Fatal("DATABASE_URL not set")
	}
	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	pgDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	// TODO: Migrate to config
	pgDB.SetMaxIdleConns(1)
	pgDB.SetMaxOpenConns(1)

	// Migrate and seed AccountType
	err = MigrateAccountType(db)
	if err != nil {
		return nil, err
	}
	err = MigrateCurrency(db)

	if err != nil {
		return nil, err
	}
	// Migrate and seed Currency
	err = MigrateCurrency(db)
	if err != nil {
		return nil, err
	}

	// Auto migrate models
	err = db.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.GroupMember{},
		&models.Expense{},
		&models.Split{},
		&models.Settlement{},
		&models.Comment{},
		&models.Account{},
	)

	if err != nil {
		return nil, err
	}

	return &GormDatabase{DB: db}, nil
}
