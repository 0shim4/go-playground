package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func test(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

func main() {
	router := gin.Default()
	router.GET("/", test)
	router.Run("localhost:8080")
}