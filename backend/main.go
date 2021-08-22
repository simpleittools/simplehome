package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/simpleittools/simplehome/routes"
	"log"
)

func routing(app *fiber.App) {
	app.Get("/", routes.Index)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := ":3030"
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routing(app)

	log.Fatal(app.Listen(PORT))
}
