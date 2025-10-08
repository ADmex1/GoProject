package controllers

import (
	"github.com/ADMex1/GoProject/models"
	"github.com/ADMex1/GoProject/services"
	"github.com/ADMex1/GoProject/utils"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadReq(ctx, "Failed to Parse Data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadReq(ctx, "Registration Failed!", err.Error())
	}
	return utils.Success(ctx, "Registration success!", user)
}
