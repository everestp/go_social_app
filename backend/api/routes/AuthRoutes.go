package routes

import (
	"server/controllers"
	"server/validation"

	"github.com/gofiber/fiber/v2"
)


func SetupAuthRoutes(app *fiber.App){
	//auth
	app.Post("/user/signup", validation.ValidateUser, controllers.Register)
	app.Post("/user/signin", validation.ValidateUser, controllers.Login)
}