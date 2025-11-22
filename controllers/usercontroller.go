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

	Token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	RefreshToken, _ := utils.RefreshToken(user.InternalID)

	var UserRespons models.UserResponse
	_ = copier.Copy(&UserRespons, &user)

	return utils.Success(ctx, "Login Success", fiber.Map{
		"Token":         Token,
		"refresh_token": RefreshToken,
		"user":          UserRespons,
	})
	// return utils.Success(ctx, "Login Successful", fiber.Map{
	// 	"Data": user,
	// })
}
func (c *UserController) GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.service.GetByPublicID(id)
	if err != nil {
		return utils.NotFound(ctx, "User not Found", err.Error())
	}
	var UserRespons models.UserResponse
	err = copier.Copy(&UserRespons, &user)
	if err != nil {
		return utils.BadReq(ctx, "Internal Server Error", err.Error())
	}

	return utils.Success(ctx, "Data Found!", UserRespons)
}
