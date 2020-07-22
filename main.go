package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	port = os.Getenv("APP_PORT")
)

// GET /sample
func getSample(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/sample", getSample)
	e.GET("/sample/", getSample)

	e.Logger.Info(e.Start(":" + port))
}
