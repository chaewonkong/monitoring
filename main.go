package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// echo prometheus
	e.Use(echoprometheus.NewMiddleware("app"))
	e.GET("/metrics", echoprometheus.NewHandler())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/slow", func(c echo.Context) error {
		time.Sleep(1 * time.Second)
		return c.String(http.StatusOK, "ok")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
