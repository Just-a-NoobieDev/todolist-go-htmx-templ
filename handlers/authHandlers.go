package handlers

import (
	"github.com/Just-A-NoobieDev/go-htmx-templ-todo/views"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	component := views.LoginPage()
	return component.Render(c.Request().Context() ,c.Response().Writer)
}