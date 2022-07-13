package Services

import (
	"HotelSystem-LearnGo/Models/Requests"
	"HotelSystem-LearnGo/Models/Responses"
)

type IUserService interface {
	Create(req Requests.UserCreateRequest) Responses.UserResponse
	Update(req Requests.UserUpdateRequest) Responses.UserResponse
	Delete(uuidString string) string
	FindByEmail(email string) Responses.UserResponse
	GetAll(request Requests.GetAllRequest) []Responses.UserResponse
	AssignRolesToUser(userId string, roleIds []string) bool
	RemoveRolesFromUser(userId string, roleIds []string) bool
	AssignBuildingsToUser(userId string, buildingIds []string) bool
	RemoveBuildingsFromUser(userId string, buildingIds []string) bool
}
