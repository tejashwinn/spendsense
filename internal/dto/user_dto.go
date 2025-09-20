package dto

import "spendsense/internal/models"

// User DTOs

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserToResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func RequestToUser(req *UserRequest) models.User {
	return models.User{
		Name:  req.Name,
		Email: req.Email,
	}
}

func UpdateUserFromRequest(user *models.User, req *UserRequest) {
	user.Name = req.Name
	user.Email = req.Email
}
