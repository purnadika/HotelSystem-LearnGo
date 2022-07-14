package main

import (
	app "HotelSystem-LearnGo/App"
	controller "HotelSystem-LearnGo/Controllers"
	Database "HotelSystem-LearnGo/Database"
	helper "HotelSystem-LearnGo/Helper"
	"HotelSystem-LearnGo/Middleware"
	repositories "HotelSystem-LearnGo/Repositories"
	"HotelSystem-LearnGo/Services"

	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	Database.Connect("hotelsystem_sa:iZU63LZh6eqzDCAf@tcp(localhost:3306)/new_hotelSystem?parseTime=true")
	Database.Migrate()
	validate := validator.New()

	//Role
	roleRepository := repositories.NewRoleRepository()
	buildingRepository := repositories.NewBuildingRepository()
	userRepository := repositories.NewUserRepository()
	roleService := Services.NewRoleService(roleRepository, validate)
	buildingService := Services.NewBuildingService(buildingRepository, validate)
	userService := Services.NewUserService(userRepository, validate, roleService, buildingService)
	roleController := controller.NewRoleController(roleService)
	buildingController := controller.NewBuildingController(buildingService)
	userController := controller.NewUserController(userService)
	roleRouter := app.Initialize(roleController, buildingController, userController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: Middleware.NewAuthMiddleware(roleRouter),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
