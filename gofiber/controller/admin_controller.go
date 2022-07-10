package controller

import (
	"github.com/gofiber/fiber/v2"
	"gofiber/util"
)

func GetUsers(c *fiber.Ctx) error {
	return util.JSON(c, 200, "success", "哈哈哈")
}
