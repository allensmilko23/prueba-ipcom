# Prueba parte 1

Desarrollada en Golang como API restfull, con echo para creacion de servidor
## Instalacion
Utilice el comando 
```bash
go get
```
Posterior el comando 
```bash
go mod tidy
```

## Uso

Puede utilizar el debugger en caso de utilizar visual studio code, de locontrario corra el comando
```bash
go run main.go
```
en caso de utilizar el comando go run, identifique el puerto en el que se levanto el servidor.

El endpoint para validar la prueba es el siguiente:
http://localhost:{{Puerto}}/resumen/2019-12-01?days=9