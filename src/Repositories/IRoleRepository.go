package repositories

import (
	dto "HotelSystem-LearnGo/Models/Dto"
	"HotelSystem-LearnGo/Models/Requests"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type IRoleRepository interface {
	Create(ctx context.Context, tx *sql.Tx, role dto.RoleDto) dto.RoleDto
	Update(ctx context.Context, tx *sql.Tx, role dto.RoleDto) dto.RoleDto
	Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) (string, error)
	FindByRoleName(ctx context.Context, tx *sql.Tx, roleName string) (dto.RoleDto, error)
	GetAll(ctx context.Context, tx *sql.Tx, req Requests.GetAllRequest) []dto.RoleDto
}
