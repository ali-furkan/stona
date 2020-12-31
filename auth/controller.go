package auth

import (
	"stona/tools/logger"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Init(router fiber.Router) {
	logger.Debug("Auth Controller", "{ path: /auth } Initializing Controller")

	aRouter := router.Group("/auth")

	aRouter.Use(limiter.New(limiter.Config{
		Max:        5,
		Expiration: 10 * time.Second,
	}))

	aRouter.Post("/verify", Service().AuthMiddleware, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Successfully Verified",
			"role":    "admin",
		})
	})

}
