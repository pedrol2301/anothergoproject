package controllers

import (
	"gotodolist/initializers"
	"gotodolist/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	user, _ := c.Get("user")
	var body struct {
		Title       string
		Description string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	todo := models.ToDo{Title: body.Title, Description: body.Description, UserID: user.(models.User).ID, Finished: false}

	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetUsersTodo(c *gin.Context) {
	userlogged, _ := c.Get("user")
	var user = models.User{}
	// var todos []models.ToDo

	initializers.DB.Where("username = ?", userlogged.(models.User).Username).Preload("Todos").Find(&user)
	c.JSON(http.StatusOK, gin.H{
		"todos": user.Todos,
	})

}

func GetUserTodoById(c *gin.Context) {
	todo := c.Param("id")
	userlogged, _ := c.Get("user")
	var user = models.User{}
	// var todos []models.ToDo

	initializers.DB.Where("username = ?", userlogged.(models.User).Username).Preload("Todos", "id = ?", todo).Find(&user)

	if len(user.Todos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todos": user.Todos[0],
	})

}

func UpdateTodoById(c *gin.Context) {
	todoId := c.Param("id")
	userlogged, _ := c.Get("user")
	// var user = models.User{}
	var todo models.ToDo

	initializers.DB.First(&todo, "id = ? AND user_id = ?", todoId, userlogged.(models.User).ID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	var body struct {
		Title       string
		Description string
		Finished    bool
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	todo.Title = body.Title
	todo.Description = body.Description
	todo.Finished = body.Finished

	initializers.DB.Save(&todo)

	c.JSON(http.StatusOK, gin.H{
		"todo": &todo,
	})
}

func DeleteTodoById(c *gin.Context) {
	todoId := c.Param("id")
	userlogged, _ := c.Get("user")
	// var user = models.User{}
	var todo models.ToDo

	initializers.DB.First(&todo, "id = ? AND user_id = ?", todoId, userlogged.(models.User).ID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}

	initializers.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{})
}
