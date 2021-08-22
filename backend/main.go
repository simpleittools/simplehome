package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/simpleittools/simplehome/database"
	"github.com/simpleittools/simplehome/routes"
)

func main() {

	database.Conn()

	PORT := ":3030"
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	app.Listen(PORT)
}
