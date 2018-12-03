package main

import (
	"os"
	"time"

	"github.com/labstack/echo"
)

var podDescription = ""

func init() {
	podDescription = "Pod started at " + time.Now().UTC().Format("15:04:05")
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"hostname":    os.Getenv("HOSTNAME"),
			"message":     os.Getenv("SYSTEM_NAME"),
			"description": podDescription,
			"debug":       "test",
		})
	})

	e.Logger.Fatal(e.Start(":5000"))
}
