package database

import (
	"github.com/tejashwinn/sependsense/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDatabase struct {
	DB *gorm.DB
}

func New(connection string) (*GormDatabase, error) {
	db, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.User{})
	return &GormDatabase{
		DB: db,
	}, nil

}
