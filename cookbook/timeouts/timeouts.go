package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func handler(c echo.Context) error {
	time.Sleep(5 * time.Second)
	return c.String(http.StatusOK, "Done!")
}

func main() {
	e := echo.New()

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 2 * time.Second,
	}))

	e.GET("/", handler)

	e.Logger.Fatal(e.Start(":1323"))
}
