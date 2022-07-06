package app

import (
	controller "HotelSystem-LearnGo/Controllers/Role"
	exceptions "HotelSystem-LearnGo/Exceptions"

	"github.com/julienschmidt/httprouter"
)

func RoleRouter(roleController controller.IRoleController) *httprouter.Router {
	router := httprouter.New()
	router.POST("/api/role/create", roleController.Create)
	router.POST("/api/role/update", roleController.Update)
	router.POST("/api/role/delete/:roleId", roleController.Delete)
	router.POST("/api/role/findbyname/:roleName", roleController.FindByRoleName)
	router.POST("/api/role/getall", roleController.GetAll)

	router.PanicHandler = exceptions.ErrorHandler
	return router
}
