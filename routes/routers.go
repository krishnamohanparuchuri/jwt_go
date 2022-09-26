package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krishna/jwt_go/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.GetUser)
	app.Post("/api/logout", controllers.Logout)
}
