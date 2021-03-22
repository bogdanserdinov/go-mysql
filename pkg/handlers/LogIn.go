package handlers

import (
	"github.com/labstack/echo"
	"text/template"
)

func logIn(c echo.Context) error{
	tmpl := template.Must(template.ParseGlob("templates/index.gohtml"))
	tmpl.Execute()
}
