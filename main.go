package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

/*************************************************************************************************/
// struct um Objekte zu erstellen

type liste struct {
	// uppercase sensitive
	Id   string `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

/*************************************************************************************************/
// Array für die CSV-Liste

var todoListe = []liste{}

/*************************************************************************************************/
// Ausgabe des Arrays

func getList(c *fiber.Ctx) error {
	loadCSV()

	return c.JSON(todoListe)
}

/*************************************************************************************************/
// Funktion zum hinzufügen eines Eintrags

func addTodo(c *fiber.Ctx) error {

	newEntry := new(liste)
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

	writeCSV()
	return c.JSON(todoListe)
}

/*************************************************************************************************/
// Funktion zum Aktualisieren der Liste

func updateList(c *fiber.Ctx) error {

	id := c.Params("id")
	err := c.BodyParser((todoListe))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	for i := 0; i < len(todoListe); i++ {
		done := false
		if todoListe[i][2] == "true" {
			done = true
		}
	}

	for _, update := range todoListe {
		if update.Id == id {
			id_int, _ := strconv.Atoi(id)
			todoListe[id_int].Done = !todoListe[id_int].Done
		}
	}

	writeCSV()
	return c.JSON(todoListe)
}

/*************************************************************************************************/
// Funktion zum löschen von Einträgen

// func deleteTodo(c *fiber.Ctx) error {

// 	id := c.Params("id")

// 	for i := 0; i < len(todoListe); i++ {
// 		if todoListe[i].Id == id {
// 			todoListwe.Delete(&id)
// 		}
// 	}

// 	writeCSV()
// 	return c.JSON(todoListe)
// }

/*************************************************************************************************/
// Funktion zum laden der CSV-Daten

func loadCSV() {
	var newList []liste

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

		readList := liste{Id: record[i][0], Name: record[i][1], Done: done}
		newList = append(newList, readList)
	}

	todoListe = newList
}

/*************************************************************************************************/
// Funktion zum schreiben in die CSV

func writeCSV() {

	file, err := os.Create("data.csv")
	if err != nil {
		fmt.Println(err)
	}

	writer := csv.NewWriter(file)
	header := []string{"id", "name", "done"}
	writer.Write(header)

	for _, record := range todoListe {
		todo := []string{record.Id, record.Name, fmt.Sprint(record.Done)}
		_ = writer.Write(todo)
	}

	writer.Flush()
	file.Close()
}

/*************************************************************************************************/

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Get("/todos", getList)
	app.Post("/todos", addTodo)
	app.Put("/todos/", updateList)
	//app.Delete("/todos/:id", deleteTodo)

	app.Listen(":5000")
}
