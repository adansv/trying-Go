package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	
)

func main() {
	// intanciar echo
	e := echo.New()

	//ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HOliis todo okish")
	})

	e.Logger.Fatal(e.Start(":1323"))

}
