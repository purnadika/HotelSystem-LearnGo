package Services

import (
	"HotelSystem-LearnGo/Models/Requests"
	"HotelSystem-LearnGo/Models/Responses"
	"context"
)

type IUserService interface {
	Create(ctx context.Context, req Requests.UserCreateRequest) Responses.UserResponse
	Update(ctx context.Context, req Requests.UserUpdateRequest) Responses.UserResponse
	Delete(ctx context.Context, uuidString string) string
	FindByEmail(ctx context.Context, email string) Responses.UserResponse
	GetAll(ctx context.Context, request Requests.GetAllRequest) []Responses.UserResponse
	AssignRolesToUser(ctx context.Context, userId string, roleIds []string) bool
	RemoveRolesFromUser(ctx context.Context, userId string, roleIds []string) bool
	AssignBuildingsToUser(ctx context.Context, userId string, buildingIds []string) bool
	RemoveBuildingsFromUser(ctx context.Context, userId string, buildingIds []string) bool
}
