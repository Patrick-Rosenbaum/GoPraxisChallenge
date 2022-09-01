package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type todos struct {
	id   int64
	name string
	done bool
}

var (
	todoListe = []todos{}
)

func loadFromCSV() {
	// open file
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()
}

func getTodo(c *fiber.Ctx) error {
	loadFromCSV()

	return c.JSON(todoListe)
}

func main() {
	app := fiber.New()

	// für die Statische Dateien wie CSS, Images und JavaScript
	app.Static("/", "./")

	// gibt Hello World zurück
	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.SendFile("./indexTest.html")
	})

	/*app.Post("/todos", func(c *fiber.Ctx) error {
		return c.JSON("/list.json")
	})*/

	app.Listen(":5000")
}
