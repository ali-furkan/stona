package main

import (
	"log"

	"stona/config"
	"stona/storages"

	"github.com/gofiber/fiber/v2"
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

	storageRouter := app.Group(config.Config().RootPath)

	storage.Init(storageRouter)

	PORT := config.GetPort()

	if err := app.Listen(PORT); err != nil {
		log.Fatalln(err)
	}
}
