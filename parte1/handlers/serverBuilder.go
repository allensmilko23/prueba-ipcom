package handlers

import (
	"fmt"
	"os"

	// "os"

	"github.com/labstack/echo/v4"
)

func ServerBuilder() (*echo.Echo, string) {
	host := "0.0.0.0"
	port := os.Getenv("PORT")
	server := echo.New()
	serverAddr := fmt.Sprintf("%s:%s", host, port)
	fmt.Println(serverAddr)
	return server, serverAddr
}
