package helper

import (
	dto "HotelSystem-LearnGo/Models/Dto"
	"HotelSystem-LearnGo/Models/Responses"
)

func ToUserResponse(user dto.UserDto) Responses.UserResponse {

	return Responses.UserResponse{
		Id:        user.Id.String(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Roles:     user.Roles,
		Address:   user.Address,
		Buildings: user.Buildings,
	}
}

func ToUserListResponse(users []dto.UserDto) []Responses.UserResponse {
	var userResponses []Responses.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
