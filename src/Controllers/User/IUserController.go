package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IUserController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUserName(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AssignRolesToUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	RemoveRolesFromUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetUsersOnBuilding(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetUsersOnRole(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
