package models

import (
	"time"

	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	GroupID  uint   `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	Role     string `gorm:"default:member"`
	JoinedAt time.Time
	Group    Group `gorm:"foreignKey:GroupID"`
	User     User  `gorm:"foreignKey:UserID"`
}

// GroupMember DTOs
type GroupMemberRequest struct {
	GroupID uint   `json:"group_id"`
	UserID  uint   `json:"user_id"`
	Role    string `json:"role"`
}

type GroupMemberResponse struct {
	ID       uint   `json:"id"`
	GroupID  uint   `json:"group_id"`
	UserID   uint   `json:"user_id"`
	Role     string `json:"role"`
	JoinedAt int64  `json:"joined_at"`
}

func GroupMemberToResponse(member *GroupMember) GroupMemberResponse {
	return GroupMemberResponse{
		ID:       member.ID,
		GroupID:  member.GroupID,
		UserID:   member.UserID,
		Role:     member.Role,
		JoinedAt: member.JoinedAt.Unix(),
	}
}

func RequestToGroupMember(req *GroupMemberRequest) GroupMember {
	return GroupMember{
		GroupID: req.GroupID,
		UserID:  req.UserID,
		Role:    req.Role,
	}
}

func UpdateGroupMemberFromRequest(member *GroupMember, req *GroupMemberRequest) {
	member.Role = req.Role
}
