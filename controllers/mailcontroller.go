package controllers

import (
	"github.com/ADMex1/GoProject/models"
	"github.com/ADMex1/GoProject/services"
	"github.com/ADMex1/GoProject/utils"
	"github.com/gofiber/fiber/v2"

	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type MailController struct {
	service services.MailService
}

func NewMailController(s services.MailService) *MailController {
	return &MailController{service: s}
}

func (c *MailController) CreateMailThread(ctx *fiber.Ctx) error {
	var userID uuid.UUID
	var errs error

	mailThread := new(models.MailThread)
	user := ctx.Locals("user").(*jwt.Token)

	claims := user.Claims.(jwt.MapClaims)
	if err := ctx.BodyParser(mailThread); err != nil {
		return utils.BadReq(ctx, "Failed to read request", err.Error())
	}
	userID, errs = uuid.Parse(claims["pub_id"].(string))
	if errs != nil {
		return utils.BadReq(ctx, "Failed to read request", errs.Error())
	}
	mailThread.SenderPublicID = userID
	if err := c.service.CreateNewMailThread(mailThread); err != nil {
		return utils.BadReq(ctx, "Failed to create a Mail thread", err.Error())
	}
	return utils.Success(ctx, "Mail thread Created!", mailThread)
}

func (c *MailController) AddReceiver(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	var userIDs []string
	if err := ctx.BodyParser(&userIDs); err != nil {
		return utils.BadReq(ctx, "Failed to parse Data", err.Error())
	}
	if err := c.service.AddReceiver(publicID, userIDs); err != nil {
		return utils.BadReq(ctx, "failed to add receiver", err.Error())
	}
	return utils.Success(ctx, "receiver added", nil)
}

func (c *MailController) CreateNewMessage(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	var userID uuid.UUID
	var errs error
	mail := new(models.Mail)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if err := ctx.BodyParser(mail); err != nil {
		return utils.BadReq(ctx, "Failed to read request", err.Error())
	}
	userID, errs = uuid.Parse(claims["pub_id"].(string))
	if errs != nil {
		return utils.BadReq(ctx, "Failed to read request", errs.Error())
	}
	mail.UserPubID = userID
	if err := c.service.CreateNewMail(publicID, mail); err != nil {
		return utils.BadReq(ctx, "Failed to create mail", err.Error())
	}
	return utils.Success(ctx, "Mail Created!", mail)

}
