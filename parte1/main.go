package main

import (
	cc "prueba-ipcom/context-components"
	"prueba-ipcom/handlers"
)

func main() {
	resumen := cc.NewResumenApi()
	server, serverAddr := handlers.ServerBuilder()
	handlers.Router(server.Group("/resumen"), resumen)
	server.Logger.Fatal(server.Start(serverAddr))
}
