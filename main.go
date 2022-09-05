package main

import (
	"github.com/gofiber/fiber/v2"
)

type list struct {
	id   int64
	name string
	done bool
}

var todoListe = []list{}

func getList(c *fiber.Ctx) error {
	//loadCSV()

	//todoList := strconv.Atoi(todoListe)
	return c.SendFile("data.json")
}

func addTodo(c *fiber.Ctx) error {
	return c.SendFile("data.json")
}

func updateList(c *fiber.Ctx) error {
	return c.SendFile("data.json")
}

func deleteTodo(c *fiber.Ctx) error {
	return c.SendFile("data.json")
}

func main() {
	app := fiber.New()

	app.Get("/todos", getList)
	app.Post("/todos", addTodo)
	app.Put("/todos", updateList)
	app.Delete("/todos/:id", deleteTodo)

	app.Listen(":5000")
}
