package routes

import (
	"backend-project/controller"
	"backend-project/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
app.Post("/api/register", controller.Register)
app.Post("/api/login", controller.Login)
app.Post("/api/logout", controller.Logout)
app.Use((middleware.AuthMiddleware))
app.Post("/api/post", controller.CreatePost)
app.Get("/api/allpost", controller.AllPost)
app.Get("/api/allpost/:id", controller.DetailPost)
app.Put("/api/updatepost/:id", controller.UpdatePost)
app.Get("/api/uniquepost", controller.UniquePost)
app.Delete("/api/deletepost/:id", controller.DeletePost)
app.Post("/api/upload-image", controller.Upload)
// get image
app.Static("/api/uploads", "./uploads")
}