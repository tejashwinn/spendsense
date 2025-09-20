package dto

import "spendsense/internal/models"

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

func GroupToResponse(group *models.Group) GroupResponse {
	return GroupResponse{
		ID:        group.ID,
		Name:      group.Name,
		OwnerID:   group.OwnerID,
		CreatedAt: group.CreatedAt.Unix(),
	}
}

func RequestToGroup(req *GroupRequest) models.Group {
	return models.Group{
		Name:    req.Name,
		OwnerID: req.OwnerID,
	}
}

func UpdateGroupFromRequest(group *models.Group, req *GroupRequest) {
	group.Name = req.Name
	group.OwnerID = req.OwnerID
}
