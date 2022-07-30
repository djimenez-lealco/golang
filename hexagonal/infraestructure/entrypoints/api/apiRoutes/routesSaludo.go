package apiRoutes

import (
	"hexagonal/infraestructure/injectors"
	"net/http"

	"github.com/labstack/echo"
)

func GetRoutesSaludo(router *echo.Echo) {

	router.GET("/saludar", func(c echo.Context) error {
		return c.String(http.StatusOK, injectors.NewSaludarInjector().Hello())
	})

}
