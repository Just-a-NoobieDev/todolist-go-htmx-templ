package main

import (
	todo "github.com/Just-A-NoobieDev/go-htmx-templ-todo/handlers"
	auth "github.com/Just-A-NoobieDev/go-htmx-templ-todo/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)




func main() {
	router := echo.New()
	router.Static("/dist", "dist")
	router.Logger.SetLevel(log.DEBUG)

	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// routes
	router.GET("/", todo.GetAllTodos)
	router.GET("/todo/:id", todo.GetTodoById)
	router.GET("/todos", todo.GetAllTodosTotal)
	router.GET("/dlt/:id", todo.DeleteBtn)
	router.GET("/refresh", todo.RefreshTodoList)
	router.GET("/refresh/:id", todo.RefreshSingleTodo)
	router.POST("/search", todo.SearchTodo)
	router.POST("/add", todo.CreateTodo)
	router.GET("/edit", todo.EditTodo)
	router.GET("/edit-status", todo.UpdateTodoStatus)
	router.PUT("/edit-description", todo.UpdateTodoDesc)
	router.DELETE("/delete/:id", todo.DeleteTodoById)

	router.GET("/login", auth.Login)

	router.Logger.Fatal(router.Start(":8080"))
}