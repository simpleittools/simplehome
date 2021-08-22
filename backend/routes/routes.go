package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/simpleittools/simplehome/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Index)

}
