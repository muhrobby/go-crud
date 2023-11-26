package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-crud/handler"
)

func Routes(app *fiber.App) {

	user := app.Group("api")
	user.Get("/users", handler.ShowUser)
	user.Get("/user/:uuid", handler.ShowByIdUser)
	user.Post("/usercreated", handler.CreateUser)
	user.Patch("/user/:uuid", handler.UpdateUser)
	user.Delete("/user/:uuid", handler.DeleteUser)

}
