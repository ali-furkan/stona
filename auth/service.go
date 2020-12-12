package auth

import (
	"encoding/base64"
	"log"
	"stona/tools/messages"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	token string
}

var authService = new(AuthService)

func Service() *AuthService {
	return authService
}

func (c *AuthService) SetToken(token string) {
	c.token = token
	return
}

func (c *AuthService) GenerateToken(key string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(key), 16)
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(hash)
}

func (c *AuthService) VerifyToken(token string, bearer bool) (val bool) {
	if bearer == true {
		val = "Bearer "+c.token == token
	} else {
		val = c.token == token
	}
	return
}

func (s *AuthService) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if !s.VerifyToken(token, true) {
		return messages.ErrorMessage(c, fiber.StatusUnauthorized, "Unauthorized request")
	}

	return c.Next()

}
