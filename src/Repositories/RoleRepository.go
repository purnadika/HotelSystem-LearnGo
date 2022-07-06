package repositories

import (
	helper "HotelSystem-LearnGo/Helper"
	dto "HotelSystem-LearnGo/Models/Dto"
	"HotelSystem-LearnGo/Models/Requests"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type RoleRepository struct {
}

func NewRoleRepository() IRoleRepository {
	return &RoleRepository{}
}

func (repo RoleRepository) Create(ctx context.Context, tx *sql.Tx, role dto.RoleDto) dto.RoleDto {
	generatedId, err := uuid.NewUUID()
	helper.PanicIfError(err)

	query := "INSERT INTO role (Id, Name, Description) VALUES (?,?,?)"
	result, err := tx.ExecContext(ctx, query, generatedId, role.Name, role.Description)
	helper.PanicIfError(err)
	result.RowsAffected()
	role.Id = generatedId
	return role
}

func (repo RoleRepository) Update(ctx context.Context, tx *sql.Tx, role dto.RoleDto) dto.RoleDto {
	query := "UPDATE role SET Name = ?, Description = ? WHERE Id = ?"
	result, err := tx.ExecContext(ctx, query, role.Name, role.Description, role.Id.String())
	helper.PanicIfError(err)
	result.RowsAffected()
	return role
}

func (repo RoleRepository) Delete(ctx context.Context, tx *sql.Tx, id uuid.UUID) (string, error) {
	query := "DELETE FROM role WHERE Id = ?"
	result, err := tx.ExecContext(ctx, query, id)
	helper.PanicIfError(err)
	totalAffected, err := result.RowsAffected()
	helper.PanicIfError(err)
	if totalAffected == 0 {
		return "", errors.New("No data exists for that ID : " + id.String())
	} else {
		return "Role deleted successfully", nil
	}
}

func (repo RoleRepository) FindByRoleName(ctx context.Context, tx *sql.Tx, roleName string) (dto.RoleDto, error) {
	query := "SELECT * FROM role WHERE Name = ?"
	result, err := tx.QueryContext(ctx, query, roleName)
	helper.PanicIfError(err)
	defer result.Close()

	role := dto.RoleDto{}
	if result.Next() {
		err := result.Scan(&role.Id, &role.Name, &role.Description)
		helper.PanicIfError(err)
		return role, nil
	} else {
		return role, errors.New("ROLE NOT FOUND")
	}
}

func (repo RoleRepository) GetAll(ctx context.Context, tx *sql.Tx, req Requests.GetAllRequest) []dto.RoleDto {
	fmt.Println(req.IsDescending)
	query := "SELECT * FROM role ORDER BY ? ASC LIMIT ?,?"
	if req.IsDescending {
		query = "SELECT * FROM role ORDER BY ? DESC LIMIT ?,?"
	}
	result, err := tx.QueryContext(ctx, query, req.OrderBy, req.Skip, req.Take)
	helper.PanicIfError(err)
	defer result.Close()

	var roles []dto.RoleDto
	for result.Next() {
		role := dto.RoleDto{}
		err := result.Scan(&role.Id, &role.Name, &role.Description)
		fmt.Println(role.Name)
		helper.PanicIfError(err)
		roles = append(roles, role)
	}
	return roles
}
