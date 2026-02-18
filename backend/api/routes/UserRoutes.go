package routes

import (


	"github.com/gofiber/fiber/v2"
	"server/controllers"
)


func SetupUserRoutes(app *fiber.App){
	//autuser
	app.Get("/user/getUser/:id",  controllers.GetUserByID)

}