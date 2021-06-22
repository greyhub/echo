package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getJSONP(c echo.Context) error {
	callback := c.QueryParam("callback")
	var content struct {
		Response  string    `json:"response"`
		Timestamp time.Time `json:"timestamp"`
		Random    int       `json:"random"`
	}
	content.Response = "Sent via JSONP"
	content.Timestamp = time.Now().UTC()
	content.Random = rand.Intn(1000)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)
	return c.JSONP(http.StatusOK, callback, &content)
}

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())

	e.Static("/", "ui")

	// JSONP
	e.GET("/jsonp", getJSONP)

	e.Logger.Fatal(e.Start(":1323"))
}
