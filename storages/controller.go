package storages

import (
	"stona/auth"

	"github.com/gofiber/fiber/v2"
)

func (m *StorageModule) Init(router fiber.Router) {
	sRouter := router.Group(m.Path)
	sRouter.Get("/:path/:id", Service().GetAsset)

	sRouter.Get("/:path", Service().GetList)

	sRouter.Put("/:path/:id", auth.Service().AuthMiddleware, Service().Upload)
	sRouter.Put("/:path", auth.Service().AuthMiddleware, Service().Upload)
	sRouter.Delete("/:path/:id", auth.Service().AuthMiddleware, Service().Delete)
}
