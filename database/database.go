package database

import (
	"log"
	"spendsense/config"
	"spendsense/internal/models"

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

	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	return &GormDatabase{DB: db}, nil
}
