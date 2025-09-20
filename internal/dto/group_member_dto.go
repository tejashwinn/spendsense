package dto

import "spendsense/internal/models"

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

func GroupMemberToResponse(member *models.GroupMember) GroupMemberResponse {
	return GroupMemberResponse{
		ID:       member.ID,
		GroupID:  member.GroupID,
		UserID:   member.UserID,
		Role:     member.Role,
		JoinedAt: member.JoinedAt.Unix(),
	}
}

func RequestToGroupMember(req *GroupMemberRequest) models.GroupMember {
	return models.GroupMember{
		GroupID: req.GroupID,
		UserID:  req.UserID,
		Role:    req.Role,
	}
}

func UpdateGroupMemberFromRequest(member *models.GroupMember, req *GroupMemberRequest) {
	member.Role = req.Role
}
