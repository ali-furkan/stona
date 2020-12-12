package messages

import "github.com/gofiber/fiber/v2"

func ErrorMessage(c *fiber.Ctx, status int, msg string) error {
	if status == 0 {
		status = 400
	}
	if msg == "" {
		msg = "Bad Request"
	}

	return c.Status(status).JSON(fiber.Map{
		"statusCode": status,
		"message":    msg,
	})
}
