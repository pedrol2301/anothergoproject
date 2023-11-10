package main

import (
	"gotodolist/controllers"
	"gotodolist/initializers"
	"gotodolist/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.POST("/post", middleware.RequireAuth, controllers.CreateTodo)
	r.GET("/post", middleware.RequireAuth, controllers.GetUsersTodo)
	r.GET("/post/:id", middleware.RequireAuth, controllers.GetUserTodoById)
	r.PUT("/post/:id", middleware.RequireAuth, controllers.UpdateTodoById)
	r.DELETE("/post/:id", middleware.RequireAuth, controllers.DeleteTodoById)

	r.Run()
}
