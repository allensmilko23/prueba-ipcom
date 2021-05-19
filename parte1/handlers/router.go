package handlers

import (
	"fmt"

	"prueba-ipcom/entities"

	"github.com/labstack/echo/v4"
)

type ResponseRequest struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"error_code"`
}

func Router(router *echo.Group, entitiesHandler []entities.RouteHandler) {
	// return c.HTML(http.StatusOK, t)

	for _, entity := range entitiesHandler {
		if entity.Method == "GET" {
			router.GET(entity.Path, entity.Function)
		} else if entity.Method == "POST" {
			router.GET(entity.Path, entity.Function)
		} else {
			response := &ResponseRequest{
				Message:   "Method not allowed",
				ErrorCode: 502,
			}
			fmt.Print(response)
			panic(response)
		}

	}
}
