package helper

import (
	Entity "HotelSystem-LearnGo/Entities"
	"HotelSystem-LearnGo/Models/Responses"
)

func ToUserResponse(user Entity.User) Responses.UserResponse {

	return Responses.UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Roles:     user.Roles,
		Address:   user.Address,
		Buildings: user.Buildings,
	}
}

func ToUserListResponse(users []Entity.User) []Responses.UserResponse {
	var userResponses []Responses.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
