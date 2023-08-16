package routers

import (
	"lebrancconvas/nheenote/models"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/", models.Init)
	app.Get("/debt", models.GetDebts)
}
