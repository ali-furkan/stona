package storages

import (
	"stona/auth"
	"stona/tools/logger"

	"github.com/gofiber/fiber/v2"
)

func (m *StorageModule) Init(router fiber.Router) {
	logger.Debug("Storage Controller", "{ path: "+m.Path+" } Initializing Controller")

	sRouter := router.Group(m.Path)
	sRouter.Get("/:path/:id", Service().GetAsset)

	sRouter.Get("/:path", auth.Service().AuthMiddleware, Service().GetList)

	sRouter.Put("/:path/:id", auth.Service().AuthMiddleware, Service().Upload)
	sRouter.Put("/:path", auth.Service().AuthMiddleware, Service().Upload)
	sRouter.Delete("/:path/:id", auth.Service().AuthMiddleware, Service().Delete)

	logger.Debug("Storage Controller", "{ path: "+m.Path+" } Initialized Controller")
}
