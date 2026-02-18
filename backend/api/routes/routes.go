package routes

import (
	"server/validaton"

	"github.com/gofiber/fiber/v2"
	"server/controllers"
)


func SetupRoutes(app *fiber.App){
	//auth
	app.Post("/user/signup", validaton.ValidateUser, controllers.Register)
	app.Post("/user/signin", validaton.ValidateUser, controllers.Login)
}