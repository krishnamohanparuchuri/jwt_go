package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krishna/jwt_go/database"
	"github.com/krishna/jwt_go/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3020")
}
