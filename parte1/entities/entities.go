package entities

import (
	"html/template"
	"io"
	"time"

	"github.com/labstack/echo/v4"
)

type RouteHandler struct {
	Path     string
	Method   string
	Function echo.HandlerFunc
}

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

type PageInfo struct {
	Title       string
	Description string
}

type ErrorHandler struct {
	ErrorCode   int32
	Description string
}

type Order struct {
	Clientid int       `json:"clientId"`
	Phone    string    `json:"phone"`
	Nombre   string    `json:"nombre"`
	Compro   bool      `json:"compro"`
	Tdc      string    `json:"tdc"`
	Monto    float64   `json:"monto"`
	Date     time.Time `json:"date"`
}

type OrdersStatistics struct {
	Total         float64 `json:"total"`
	Comprasportdc struct {
		Oro  int `json:"oro"`
		Amex int `json:"amex"`
	} `json:"comprasPorTDC"`
	Nocompraron   int     `json:"nocompraron"`
	Compramasalta float64 `json:"compraMasAlta"`
}

type RequestChan struct {
	Date string
	Day  string
}
