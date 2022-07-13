package controller

import (
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Models/Requests"
	"HotelSystem-LearnGo/Models/Responses"
	"HotelSystem-LearnGo/Services"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type RoleController struct {
	RoleService Services.IRoleService
}

func NewRoleController(roleService Services.IRoleService) IRoleController {
	return &RoleController{
		RoleService: roleService,
	}
}

func (controller *RoleController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleCreateRequest := Requests.RoleCreateRequest{}
	helper.ReadFromRequestBody(request, &roleCreateRequest)
	roleResponse := controller.RoleService.Create(roleCreateRequest)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *RoleController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleUpdateRequest := Requests.RoleUpdateRequest{}
	helper.ReadFromRequestBody(request, &roleUpdateRequest)
	roleResponse := controller.RoleService.Update(roleUpdateRequest)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *RoleController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleIdString := params.ByName("roleId")
	roleId, err := strconv.ParseUint(roleIdString, 10, 32)
	helper.PanicIfError(err)
	roleResponse := controller.RoleService.Delete(uint(roleId))
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *RoleController) FindByRoleName(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	roleName := params.ByName("roleName")
	roleResponse := controller.RoleService.FindByName(roleName)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   roleResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *RoleController) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	getAllRequest := Requests.GetAllRequest{}
	helper.ReadFromRequestBody(request, &getAllRequest)
	rolesResponse := controller.RoleService.GetAll(getAllRequest)
	jsontext, err := json.Marshal(rolesResponse)
	helper.PanicIfError(err)
	log.Println(string(jsontext))
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   rolesResponse,
	}
	helper.WriteToResponseBody(writer, response)
}
