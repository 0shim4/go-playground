package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

var books = []book {
	{ID: "1", Title: "Go API Tutorial", Author: "TechWithTim", Quantity: 2},
	{ID: "2", Title: "Make An API With Go", Author: "GoAPI", Quantity: 5},
	{ID: "3", Title: "So easy", Author: "Golang", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func test(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func main() {
	router := gin.Default()
	router.GET("/", test)
	router.GET("/books", getBooks)
	router.Run("localhost:8080")
}