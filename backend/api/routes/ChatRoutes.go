package routes

import (
	"server/controllers"


	"github.com/gofiber/fiber/v2"
)


func SetupChatRoutes(app *fiber.App){
	//auth
	app.Post("/chat/sendmessage", controllers.SendMessage)
	app.Get("/chat/getmsgsbynums", controllers.GetMsgsByNums)
	app.Get("/chat/get-user-unreadedmsg", controllers.GetUserUnreadedMsg)
	app.Get("/chat/mark-msg-asreaded", controllers.MarkMsgAsReaded)

}