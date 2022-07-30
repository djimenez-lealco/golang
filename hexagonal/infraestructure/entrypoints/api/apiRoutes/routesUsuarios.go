package apiRoutes

import (
	"hexagonal/infraestructure/injectors"
	"net/http"

	"github.com/labstack/echo"
)

func GetRoutesUsuarios(router *echo.Echo) {

	router.GET("/usuario/:id", func(c echo.Context) error {
		idUsuario := c.Param("id")
		return c.JSON(http.StatusOK, injectors.NewUsuarioInjector().GetUsuario(idUsuario))

	})
}
