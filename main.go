package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"./handler"
)

func main() {
	println("hogehoge")
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", handler.HelloHandler())

	e.Start(":1234")
}
