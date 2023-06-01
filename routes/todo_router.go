package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noergitkat/go-todo-app/app/controllers"
)

func SetupTodoRoutes(app *fiber.App) {
	app.Get("/todos", controllers.GetTodos)
	app.Post("/todos", controllers.CreateTodo)
	app.Put("/todos/:id", controllers.UpdateTodo)
	app.Delete("/todos/:id", controllers.DeleteTodo)
}
