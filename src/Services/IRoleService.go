package Services

import (
	"HotelSystem-LearnGo/Models/Requests"
	"HotelSystem-LearnGo/Models/Responses"
	"context"
)

type IRoleService interface {
	Create(ctx context.Context, req Requests.RoleCreateRequest) Responses.RoleResponse
	Update(ctx context.Context, req Requests.RoleUpdateRequest) Responses.RoleResponse
	Delete(ctx context.Context, uuidString string) string
	FindByRoleName(ctx context.Context, roleName string) Responses.RoleResponse
	GetAll(ctx context.Context, request Requests.GetAllRequest) []Responses.RoleResponse
}
