package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type list struct {
	// uppercase sensitive
	Id   string
	Name string
	Done bool
}

var todoListe = []list{}

func getList(c *fiber.Ctx) error {
	loadCSV()

	return c.JSON(todoListe)
}

func addTodo(c *fiber.Ctx) error {

	newEntry := new(list)
	err := c.BodyParser((newEntry))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	//max := 0
	for i := 0; i < len(todoListe); i++ {
		currId := todoListe[i]
		fmt.Println(currId)
		/*currId = strconv.Atoi(todoListe[i]["Id"])
		if currId > max {
			max = currId
		}*/
	}

	return c.JSON(todoListe)
}

func updateList(c *fiber.Ctx) error {

	var id = c.BodyParser("Id")

	for i := 0; i < len(todoListe); i++ {
		if todoListe[i]["Id"] == id {
			todoListe[i]["Done"] = c.BodyParser()["Done"]
		}
	}

	return c.JSON(todoListe)
}

func deleteTodo(c *fiber.Ctx) error {

	id := c.Params("Id")

	for i := 0; i < len(todoListe); i++ {
		if todoListe[i]["Id"] == id {

		}
	}

	return c.JSON(todoListe)
}

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

	app.Get("/todos", getList)
	app.Post("/todos", addTodo)
	app.Put("/todos", updateList)
	app.Delete("/todos/:id", deleteTodo)

	app.Listen(":5000")
}
