package models

import (
	"lebrancconvas/nheenote/db"
	"lebrancconvas/nheenote/forms"

	"github.com/gofiber/fiber/v2"
)

func Init(c *fiber.Ctx) error {
	return c.JSON("Database Connected!")
}

func GetDebts(c *fiber.Ctx) error {
	db := db.DBConn

	var debts []forms.Dept

	db.Find(&debts)

	return c.JSON(&debts)
}
