package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thvinhtruong/legoha/app/apperror"
)

func Response(c *fiber.Ctx, httpCode int, ReceivedError error, data interface{}) error {
	var msg string
	if ReceivedError != nil {
		msg = ReceivedError.Error()
	} else {
		msg = "success"
	}

	field := getErrorField(ReceivedError)
	if field != "" {
		msg = "success"
	}
	code := apperror.GetCode(ReceivedError)

	c.Set("Content-Type", "application/json")
	return c.Status(httpCode).JSON(fiber.Map{
		"data": data,
		"error": ErrorResponse{
			ErrorCode:  code,
			ErrorMsg:   msg,
			ErrorField: field,
		},
	})
}
