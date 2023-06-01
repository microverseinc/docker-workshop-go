package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/noergitkat/go-todo-app/app/controllers"
	"github.com/noergitkat/go-todo-app/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./public", ".html"),
	})

	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database!")

	err = db.AutoMigrate(&controllers.Todo{})
	if err != nil {
		log.Fatal(err)
	}

	controllers.InitDatabase(db)

	routes.SetupTodoRoutes(app)
	app.Static("/", "./public")

	app.Listen(":3000")
}
