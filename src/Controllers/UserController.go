package controller

import (
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Models/Requests"
	"HotelSystem-LearnGo/Models/Responses"
	"HotelSystem-LearnGo/Services"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	UserService Services.IUserService
}

func NewUserController(userService Services.IUserService) IUserController {
	return &UserController{
		UserService: userService,
	}
}

func (controller *UserController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := Requests.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)
	userResponse := controller.UserService.Create(userCreateRequest)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *UserController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := Requests.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)
	userResponse := controller.UserService.Update(userUpdateRequest)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *UserController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userIdString := params.ByName("userId")
	userId := helper.StringToUint(userIdString)
	userResponse := controller.UserService.Delete(uint(userId))
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *UserController) FindByUsername(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	username := params.ByName("username")
	userResponse := controller.UserService.FindByUsername(username)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *UserController) FindByEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	username := params.ByName("email")
	userResponse := controller.UserService.FindByEmail(username)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *UserController) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	getAllRequest := Requests.GetAllRequest{}
	helper.ReadFromRequestBody(request, &getAllRequest)
	usersResponse := controller.UserService.GetAll(getAllRequest)
	jsontext, err := json.Marshal(usersResponse)
	helper.PanicIfError(err)
	log.Println(string(jsontext))
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   usersResponse,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *UserController) AssignRolesToUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	req := Requests.UserRolesRequest{}
	helper.ReadFromRequestBody(request, &req)
	result := controller.UserService.AssignRolesToUser(req.UserId, req.RoleIds)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *UserController) AssignBuildingsToUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	req := Requests.UserBuildingsRequest{}
	helper.ReadFromRequestBody(request, &req)
	result := controller.UserService.AssignBuildingsToUser(req.UserId, req.BuildingIds)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *UserController) RemoveRolesFromUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	req := Requests.UserRolesRequest{}
	helper.ReadFromRequestBody(request, &req)
	result := controller.UserService.RemoveRolesFromUser(req.UserId, req.RoleIds)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteToResponseBody(writer, response)
}

func (controller *UserController) RemoveBuildingsFromUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	req := Requests.UserBuildingsRequest{}
	helper.ReadFromRequestBody(request, &req)
	result := controller.UserService.RemoveBuildingsFromUser(req.UserId, req.BuildingIds)
	response := Responses.GeneralResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteToResponseBody(writer, response)
}
