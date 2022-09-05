package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type todos struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
}

var (
	jsonListe = []todos{}
	todoListe = []todos{}
)

func loadFromCSV(source, destination string) error {
	// open file
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer jsonFile.Close()

	// var jsonListe []todos
	if err := json.NewDecoder(jsonFile).Decode(&jsonListe); err != nil {
		return err
	}

	outputFile, err := os.Create("data.csv")
	if err != nil {
		return err
	}

	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	writer.Flush()

	header := []string{"id", "name", "done"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, r := range jsonListe {
		var csvRow []string
		csvRow = append(csvRow, fmt.Sprint(r.Id, r.Done), r.Name)
		if err := writer.Write(csvRow); err != nil {
			return err
		}
	}
	return nil
}

func getTodo(c *fiber.Ctx) error {
	if err := loadFromCSV("data.json", "data.csv"); err != nil {
		log.Fatal(err)
	}

	return c.JSON(jsonListe)
}

func main() {
	app := fiber.New()

	app.Get("/todos", getTodo)

	app.Post("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todoListe)
	})

	app.Listen(":5000")
}
