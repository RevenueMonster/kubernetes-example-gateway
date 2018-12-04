package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

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
		return c.String(200, os.Getenv("SECRET_KEY"))
	})

	e.Logger.Fatal(e.Start(":5000"))
}
