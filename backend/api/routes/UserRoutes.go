package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gofiber/fiber/v2"
)


func SetupUserRoutes(app *fiber.App){
	//autuser
	app.Get("/user/getUser/:id",  controllers.GetUserByID)

	//Update
	app.Patch("/user/Update/:id", middleware.AuthMiddleware , controllers.UpdateUser)
}