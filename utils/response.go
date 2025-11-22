package utils

import "github.com/gofiber/fiber/v2"

type response struct {
	Status       string      `json:"status"`
	ResponseCode int         `json:"ResponseCode"`
	Message      string      `json:"message,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Error        string      `json:"error"`
}

func Success(c *fiber.Ctx, Message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(response{
		Status:       "success",
		ResponseCode: fiber.StatusOK,
		Message:      Message,
		Data:         data,
	})
}
func Created(c *fiber.Ctx, Message string) error {
	return c.Status(fiber.StatusCreated).JSON(response{
		Status:       "Created",
		ResponseCode: fiber.StatusCreated,
		Message:      Message,
	})
}
func BadReq(c *fiber.Ctx, Message string, err string) error {
	return c.Status(fiber.StatusBadRequest).JSON(response{
		Status:       "Bad Request",
		ResponseCode: fiber.StatusBadRequest,
		Message:      Message,
		Error:        err,
	})
}
func NotFound(c *fiber.Ctx, Message string, err string) error {
	return c.Status(fiber.StatusNotFound).JSON(response{
		Status:       "Not Found",
		ResponseCode: fiber.StatusNotFound,
		Message:      Message,
		Error:        err,
	})
}
func UnauthorizedAccess(c *fiber.Ctx, Message string, err string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(response{
		Status:       "Not Found",
		ResponseCode: fiber.StatusUnauthorized,
		Message:      Message,
		Error:        err,
	})
}
