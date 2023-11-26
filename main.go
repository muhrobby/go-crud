package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhrobby/go-crud/database"
	"github.com/muhrobby/go-crud/migration"
	"github.com/muhrobby/go-crud/router"
)

func main() {
	app := fiber.New()

	database.ConnectDB()
	migration.AutoMigrate()

	router.Routes(app)

	app.Listen(":3030")
}
