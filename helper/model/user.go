package model

import (
	"github.com/nostracode/mangrove-api/model/domain"
	userweb "github.com/nostracode/mangrove-api/model/web/user_web"
)

func ToUserResponse(user *domain.User) *userweb.UserResponse {

	return &userweb.UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		RoleId:      user.RoleId,
		ImageUrl:    user.ImageUrl,
		FirstVisit:  user.FirstVisit,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

}

func ToUserResponses(users []*domain.User) []*userweb.UserResponse {

	userResponses := []*userweb.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses

}
