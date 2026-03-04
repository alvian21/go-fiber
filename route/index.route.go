package route

import (
	"go-fiber-gorm/config"
	"go-fiber-gorm/handler"
	"go-fiber-gorm/middleware"
	"go-fiber-gorm/utils"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	r.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Application is running",
			"status":  "ok",
		})
	})

	r.Static("/public", config.ProjectRootPath+"/public/asset")

	r.Post("/login", handler.LoginHandler)

	user := r.Group("/user", middleware.Auth)
	user.Get("/", handler.UserHandlerGetAll)
	user.Get("/:id", handler.UserHandlerGetById)
	user.Post("/", handler.UserHandlerCreate)
	user.Put("/:id", handler.UserHandlerUpdate)
	user.Delete("/:id", handler.UserHandlerDelete)

	r.Post("/book", utils.HandleSingleFile, handler.BookHandlerCreate)

	r.Post("/gallery", utils.HandleMultipleFile, handler.PhotoHandlerCreate)
	r.Delete("/gallery/:id", handler.PhotoHandlerDelete)
}
