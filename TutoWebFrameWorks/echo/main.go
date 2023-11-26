package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", handler)
	e.GET("/user/:userId", handlerUser)

	fmt.Println("hello echo")

	e.Use(middleware.Logger())

	e.Start(":8080")
}

func handler(c echo.Context) error {
	return c.String(http.StatusOK, "Hola ECHO")
}

func handlerUser(c echo.Context) error {
	userId := c.Param("userId")
	return c.String(http.StatusOK, userId)
}
