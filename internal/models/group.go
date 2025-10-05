package models

import (
	"time"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	OwnerID   uint   `gorm:"not null"`
	Owner     User   `gorm:"foreignKey:OwnerID"`
	CreatedAt time.Time
}

// Group DTOs
type GroupRequest struct {
	Name    string `json:"name"`
	OwnerID uint   `json:"owner_id"`
}

type GroupResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	OwnerID   uint   `json:"owner_id"`
	CreatedAt int64  `json:"created_at"`
}

func GroupToResponse(group *Group) GroupResponse {
	return GroupResponse{
		ID:        group.ID,
		Name:      group.Name,
		OwnerID:   group.OwnerID,
		CreatedAt: group.CreatedAt.Unix(),
	}
}

func RequestToGroup(req *GroupRequest) Group {
	return Group{
		Name:    req.Name,
		OwnerID: req.OwnerID,
	}
}

func UpdateGroupFromRequest(group *Group, req *GroupRequest) {
	group.Name = req.Name
	group.OwnerID = req.OwnerID
}
