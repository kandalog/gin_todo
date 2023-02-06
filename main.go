package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var todos = []todo{
	{ID: 1, Title: "タイトルA", Author: "Taro"},
	{ID: 2, Title: "タイトルB", Author: "Jiro"},
	{ID: 3, Title: "タイトルC", Author: "Takuya"},
}

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)

	// ":8080"
	router.Run()
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}
