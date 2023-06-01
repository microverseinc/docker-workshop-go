package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDatabase initializes the database connection
func InitDatabase(database *gorm.DB) {
	db = database
}

// Todo represents a todo item
type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// GetTodos retrieves all todos from the database
func GetTodos(c *fiber.Ctx) error {
	var todos []Todo
	db.Find(&todos)
	return c.JSON(todos)
}


// CreateTodo creates a new todo and stores it in the database
func CreateTodo(c *fiber.Ctx) error {
	var todos []Todo
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	db.Create(&todo).Find((&todos))
	return c.JSON(todos)
}

// UpdateTodo updates a todo in the database
func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	db.Model(&Todo{}).Where("id = ?", id).Updates(Todo{Title: todo.Title, Completed: todo.Completed})
	return c.JSON(todo)
}

// DeleteTodo
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db.Delete(&Todo{}, id)
	
	var todos []Todo
	db.Find(&todos)
	return c.JSON(todos)
}