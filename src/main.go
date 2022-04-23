package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
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

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
	for i, book := range books {
		if book.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book no found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func test(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func main() {
	router := gin.Default()
	router.GET("/", test)
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("books", createBook)
	router.Run("localhost:8080")
}