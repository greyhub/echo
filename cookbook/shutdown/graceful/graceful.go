package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tylerb/graceful"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Gracefull OK")
	})
	e.Server.Addr = ":1323"

	// Serve it like a boss
	graceful.ListenAndServe(e.Server, 5*time.Second)
}
