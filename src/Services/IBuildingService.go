package Services

import (
	"HotelSystem-LearnGo/Models/Requests"
	"HotelSystem-LearnGo/Models/Responses"
	"context"
)

type IBuildingService interface {
	Create(ctx context.Context, req Requests.BuildingCreateRequest) Responses.BuildingResponse
	Update(ctx context.Context, req Requests.BuildingUpdateRequest) Responses.BuildingResponse
	Delete(ctx context.Context, uuidString string) string
	FindByBuildingName(ctx context.Context, buildingName string) Responses.BuildingResponse
	GetAll(ctx context.Context, request Requests.GetAllRequest) []Responses.BuildingResponse
}
