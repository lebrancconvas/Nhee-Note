package main

import (
	"lebrancconvas/nheenote/db"
	"lebrancconvas/nheenote/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.InitDB()
	routers.SetupRouter(app)
	defer db.CloseDB()

	app.Listen(":8001")
}
