package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "grey"
	// cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.MaxAge = 10
	c.SetCookie(cookie)

	cookie = new(http.Cookie)
	cookie.Name = "color"
	cookie.Value = "blue"
	cookie.MaxAge = 10
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

// func readCookie(c echo.Context) error {
// 	cookie, err := c.Cookie("username")
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(cookie.Name)
// 	fmt.Println(cookie.Value)
// 	return c.String(http.StatusOK, "read a cookie")
// }

func readAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK, "read all the cookies")
}

func main() {
	e := echo.New()

	e.POST("/", writeCookie)
	e.GET("/", readAllCookies)

	e.Logger.Fatal(e.Start(":1323"))
}
