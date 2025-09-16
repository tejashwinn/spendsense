package database

import (
	"github.com/gin-gonic/gin"
	"github.com/tejashwinn/sependsense/internal/model"
	"gorm.io/gorm"
)

func (db *GormDatabase) CreateUser(ctx *gin.Context, user *model.User) error {
	result := gorm.WithResult()
	err := gorm.G[model.User](db.DB, result).Create(ctx, user)
	return err
}

func (db *GormDatabase) GetById(ctx *gin.Context, id uint64) (*model.User, error) {
	user, err := gorm.G[model.User](db.DB).Where("id = ?", id).First(ctx)
	return &user, err
}
