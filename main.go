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
	router.POST("/todos", postTodo)

	// ":8080"
	router.Run()
}

// 全てのTodoを取得
func getTodos(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, todos)
}

// 新規投稿
func postTodo(c *gin.Context) {
	// bodyの内容が入る
	var newTodo todo

	// JSONを構造体にbindする
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}
