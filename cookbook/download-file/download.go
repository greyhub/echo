package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func home(c echo.Context) error {
	return c.File("index.html")
}

func show(c echo.Context) error {
	return c.File("echo.svg")
}

func inline(c echo.Context) error {
	return c.Inline("inline.txt", "Inline Text")
}

func attachment(c echo.Context) error {
	return c.Attachment("attachment.txt", "Attachment Text")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", home)
	e.GET("/file", show)
	e.GET("/inline", inline)
	e.GET("/attachment", attachment)

	e.Logger.Fatal(e.Start(":1323"))
}
