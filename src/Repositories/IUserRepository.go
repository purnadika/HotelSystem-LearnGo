package repositories

import (
	dto "HotelSystem-LearnGo/Models/Dto"
	"HotelSystem-LearnGo/Models/Requests"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type IUserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, role dto.UserDto) dto.UserDto
	Update(ctx context.Context, tx *sql.Tx, role dto.UserDto) dto.UserDto
	Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) (string, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (dto.UserDto, error)
	GetAll(ctx context.Context, tx *sql.Tx, req Requests.GetAllRequest) []dto.UserDto
}
