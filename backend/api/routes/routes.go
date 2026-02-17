package routes

import (
	"server/validaton"

	"github.com/gofiber/fiber/v2"
	"server/controller"
)


func SetupRoutes(app *fiber.App){
	//auth
	app.Post("/user/signup", validaton.ValidateUser, controller.Register)
}