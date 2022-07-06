package repositories

import (
	dto "HotelSystem-LearnGo/Models/Dto"
	"HotelSystem-LearnGo/Models/Requests"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type IBuildingRepository interface {
	Create(ctx context.Context, tx *sql.Tx, role dto.BuildingDto) dto.BuildingDto
	Update(ctx context.Context, tx *sql.Tx, role dto.BuildingDto) dto.BuildingDto
	Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) (string, error)
	FindByBuildingName(ctx context.Context, tx *sql.Tx, buildingName string) (dto.BuildingDto, error)
	GetAll(ctx context.Context, tx *sql.Tx, req Requests.GetAllRequest) []dto.BuildingDto
}
