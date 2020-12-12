package main

import (
	"log"

	"stona/config"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	PORT := config.GetPort()

	if err := app.Listen(PORT); err != nil {
		log.Fatalln(err)
	}
}
