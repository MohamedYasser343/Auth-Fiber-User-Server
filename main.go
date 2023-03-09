package main

import (
	"github.com/MohamedYasser343/database"
	"github.com/MohamedYasser343/database/migration"
	"github.com/MohamedYasser343/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseInit()
	migration.RunMigration()
	app := fiber.New()
	routes.RouteInit(app)
	app.Listen(":3000")
}
