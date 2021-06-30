package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func home(c echo.Context) (err error) {
	pusher, ok := c.Response().Writer.(http.Pusher)
	if ok {
		if err = pusher.Push("/app.css", nil); err != nil {
			return
		}
		if err = pusher.Push("/app.js", nil); err != nil {
			return
		}
		if err = pusher.Push("/echo.png", nil); err != nil {
			return
		}
	}
	return c.File("index.html")
}

func main() {
	e := echo.New()

	e.Static("/", "static")
	e.GET("/", home)

	e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
}
