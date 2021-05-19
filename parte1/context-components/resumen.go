package contextComponents

import (
	apiControllers "prueba-ipcom/controllers"
	"prueba-ipcom/entities"
)

func NewResumenApi() []entities.RouteHandler {
	var routerResumen []entities.RouteHandler

	var resumen entities.RouteHandler
	resumen.Path = "/:date"
	resumen.Method = "GET"
	resumen.Function = apiControllers.GetResumen

	routerResumen = append(routerResumen, resumen)

	return routerResumen
}
