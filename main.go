package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":5000") // listen and serve on 0.0.0.0:8080

	/*r.POST("/todos", func(c *gin.Context) {
		lastId := 0
		for i := 0; i < len(todoListe); i++ {
			currentId := parseInt(todoListe[i]["id"])

			if currentId > lastId {
				lastId = currentId
			}
		}
	})*/

	rGroup := r.Group("/todos")
	rGroup.GET("")
	rGroup.POST("")
	rGroup.PUT("")
	rGroup.DELETE("")
}
