package Services

import (
	exceptions "HotelSystem-LearnGo/Exceptions"
	helper "HotelSystem-LearnGo/Helper"
	dto "HotelSystem-LearnGo/Models/Dto"
	"HotelSystem-LearnGo/Models/Requests"
	"HotelSystem-LearnGo/Models/Responses"
	repositories "HotelSystem-LearnGo/Repositories"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type RoleService struct {
	RoleRepository repositories.IRoleRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewRoleService(repo repositories.IRoleRepository, DB *sql.DB, validate *validator.Validate) IRoleService {
	return &RoleService{
		RoleRepository: repo,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *RoleService) Create(ctx context.Context, req Requests.RoleCreateRequest) Responses.RoleResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role := dto.RoleDto{
		Name:        req.Name,
		Description: req.Description,
	}

	role = service.RoleRepository.Create(ctx, tx, role)
	return helper.ToRoleResponse(role)
}

func (service *RoleService) Update(ctx context.Context, req Requests.RoleUpdateRequest) Responses.RoleResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roleId, err := uuid.Parse(req.Id)
	helper.PanicIfError(err)

	role := dto.RoleDto{
		Id:          roleId,
		Name:        req.Name,
		Description: req.Description,
	}

	role = service.RoleRepository.Update(ctx, tx, role)
	return helper.ToRoleResponse(role)
}

func (service *RoleService) Delete(ctx context.Context, id string) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	roleId, err := uuid.Parse(id)
	helper.PanicIfError(err)
	message, err := service.RoleRepository.Delete(ctx, tx, roleId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	return message
}

func (service *RoleService) FindByRoleName(ctx context.Context, name string) Responses.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	role, err := service.RoleRepository.FindByRoleName(ctx, tx, name)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	return helper.ToRoleResponse(role)
}

func (service *RoleService) GetAll(ctx context.Context, req Requests.GetAllRequest) []Responses.RoleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}
	return helper.ToListResponse(service.RoleRepository.GetAll(ctx, tx, req))
}
