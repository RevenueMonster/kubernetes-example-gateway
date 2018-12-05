package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {

		return c.JSON(200, map[string]interface{}{
			"hostname": os.Getenv("HOSTNAME"),
			"message":  os.Getenv("SYSTEM_NAME"),
		})
	})

	e.GET("/helloworld", func(c echo.Context) error {
		response, err := http.Get("http://rm-service")
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"error": err.Error(),
			})
		}

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return c.JSON(500, map[string]interface{}{
				"error": err.Error(),
			})
		}

		return c.JSON(200, map[string]interface{}{
			"message": string(contents),
		})
	})

	e.GET("/secret", func(c echo.Context) error {

		return c.JSON(200, map[string]interface{}{
			"message": os.Getenv("SECRET_KEY"),
		})
	})

	e.GET("/configmap", func(c echo.Context) error {

		return c.JSON(200, map[string]interface{}{
			"message": os.Getenv("ENV"),
		})
	})

	e.GET("/sleep", func(c echo.Context) error {
		time.Sleep(10 * time.Second)

		return c.JSON(200, map[string]interface{}{
			"message": "sleep 10 second",
		})
	})

	e.GET("/panic", func(c echo.Context) error {
		log.Fatal("panic")
		return c.JSON(200, map[string]interface{}{
			"message": "exit",
		})
	})

	e.Logger.Fatal(e.Start(":5000"))
}
