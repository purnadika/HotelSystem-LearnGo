package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IUserController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUsername(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AssignRolesToUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	RemoveRolesFromUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	AssignBuildingsToUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	RemoveBuildingsFromUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
