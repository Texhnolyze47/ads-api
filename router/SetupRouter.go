package router

import (
	"comerciales-api/controller"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupRouter(e *echo.Echo) (*echo.Echo, error) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/cliente", controller.AddCliente)
	e.Logger.Fatal(e.Start(":1323"))

	return e, nil
}
