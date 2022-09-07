package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type list struct {
	// uppercase sensitive
	Id   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var todoListe = []list{}

func getList(c *fiber.Ctx) error {
	loadCSV()

	return c.JSON(todoListe)
}

func addTodo(c *fiber.Ctx) error {
	fmt.Println("test")
	newEntry := new(list)
	err := c.BodyParser((newEntry))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	lastId := 0
	for i := 0; i < len(todoListe); i++ {
		currId, _ := strconv.Atoi(todoListe[i].Id)
		if currId > lastId {
			lastId = currId
		}
	}

	newEntry.Id = strconv.Itoa(lastId + 1)
	todoListe = append(todoListe, *newEntry)
	return c.JSON(todoListe)
}

/*func updateList(c *fiber.Ctx) error {

todo := c.BodyParser(todoListe)

for i := 0; i < len(todoListe); i++ {
	if todo[i] == false {
		todoListe[i].Done = "true"
	}

}*/

/*id := c.Params("id")
	body := list{}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	update := list{}

	if newUpdate := c.todoListe.First(&update, id); update.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, update.Error.Error())
	}

	return c.JSON(todoListe)
}*/

/*func deleteTodo(c *fiber.Ctx) error {

	id := c.Params("id")

	for i := 0; i < len(todoListe); i++ {
		if todoListe[i].Id == id {
			todoListe.Delete(&id)
		}
	}

	return c.JSON(todoListe)
}*/

func loadCSV() {
	var newList []list

	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	record, _ := reader.ReadAll()

	for i := 1; i < len(record); i++ {
		done := false
		if record[i][2] == "true" {
			done = true
		}

		readList := list{Id: record[i][0], Name: record[i][1], Done: done}
		newList = append(newList, readList)
	}
	todoListe = newList
}

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Get("/todos", getList)
	app.Post("/todos", addTodo)
	//app.Put("/todos", updateList)
	//app.Delete("/todos/:id", deleteTodo)

	app.Listen(":5000")
}
