package controllers

import (
	"github.com/ADMex1/GoProject/models"
	"github.com/ADMex1/GoProject/services"
	"github.com/ADMex1/GoProject/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
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
	var UserRespons models.UserResponse
	_ = copier.Copy(&UserRespons, user)
	return utils.Success(ctx, "Registration success!", UserRespons)
}
func (l *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadReq(ctx, "Invalid request", err.Error())
	}
	user, err := l.service.Login(body.Email, body.Password)
	if err != nil {
		return utils.UnauthorizedAccess(ctx, "login failed", err.Error())
	}

	token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	refreshtoken, _ := utils.RefreshToken(user.InternalID)

	var UserRespons models.UserResponse
	_ = copier.Copy(&UserRespons, &user)

	return utils.Success(ctx, "Login Success", fiber.Map{
		"access_token":  token,
		"refresh_token": refreshtoken,
		"user":          UserRespons,
	})
}
