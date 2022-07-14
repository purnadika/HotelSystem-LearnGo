package app

import (
	controller "HotelSystem-LearnGo/Controllers"
	exceptions "HotelSystem-LearnGo/Exceptions"

	"github.com/julienschmidt/httprouter"
)

func Initialize(roleController controller.IRoleController, buildingController controller.IBuildingController,
	userController controller.IUserController) *httprouter.Router {
	router := httprouter.New()
	//Role
	router.POST("/api/role/create", roleController.Create)
	router.POST("/api/role/update", roleController.Update)
	router.POST("/api/role/delete/:roleId", roleController.Delete)
	router.POST("/api/role/findbyname/:roleName", roleController.FindByRoleName)
	router.POST("/api/role/getall", roleController.GetAll)

	//Building
	router.POST("/api/building/create", buildingController.Create)
	router.POST("/api/building/update", buildingController.Update)
	router.POST("/api/building/delete/:buildingId", buildingController.Delete)
	router.POST("/api/building/findbyname/:buildingName", buildingController.FindByBuildingName)
	router.POST("/api/building/getall", buildingController.GetAll)

	//User
	router.POST("/api/user/create", userController.Create)
	router.POST("/api/user/update", userController.Update)
	router.POST("/api/user/delete/:userid", userController.Delete)
	router.POST("/api/user/findbyusername/:username", userController.FindByUsername)
	router.POST("/api/user/findbyemail/:email", userController.FindByEmail)
	router.POST("/api/user/getall", userController.GetAll)
	router.POST("/api/user/assignroles", userController.AssignRolesToUser)
	router.POST("/api/user/removeroles", userController.RemoveRolesFromUser)
	router.POST("/api/user/assignbuildings", userController.AssignBuildingsToUser)
	router.POST("/api/user/removebuildings", userController.RemoveBuildingsFromUser)

	router.PanicHandler = exceptions.ErrorHandler
	return router
}
