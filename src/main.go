package main

import (
	app "HotelSystem-LearnGo/App"
	controller "HotelSystem-LearnGo/Controllers/Role"
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

	roleRepository := repositories.NewRoleRepository()
	roleService := Services.NewRoleService(roleRepository, validate)
	roleController := controller.NewRoleController(roleService)
	router := app.RoleRouter(roleController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: Middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
