package main

import (
	"log"
	"strings"
	"time"

	"stona/config"
	"stona/storages"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

var storage *storages.StorageModule

func init() {
	// Only one storage can add for now.
	// In the future multiple storages can
	// be created with configuration

	storage = storages.NewStorage(config.Config().StoragePath, "firebase")

}

func main() {
	app := fiber.New()

	app.Use(etag.New())
	app.Use(csrf.New())
	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(limiter.New(limiter.Config{
		Max:        24,
		Expiration: 30 * time.Second,
	}))

	format := []string{color.HiBlackString("${time}"), color.BlueString("[Request]"), color.CyanString("${status} - ${method} ${path}"), color.YellowString("+${latency}\n")}
	app.Use(logger.New(logger.Config{
		Format: strings.Join(format, " "),
	}))

	storageRouter := app.Group(config.Config().RootPath)

	storage.Init(storageRouter)

	PORT := config.GetPort()

	if err := app.Listen(PORT); err != nil {
		log.Fatalln(err)
	}
}
