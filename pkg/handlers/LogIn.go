package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func LogIn(c echo.Context) error{
	return c.Render(http.StatusOK,"././templates/index.html",nil)
}
