package main

import (
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Grace OK")
	})

	e.Server.Addr = ":1323"

	// Serve it like a boss
	e.Logger.Fatal(gracehttp.Serve(e.Server))
}
