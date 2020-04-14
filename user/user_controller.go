package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	p := &User{
		Id:   1,
		Name: "John",
	}
	return c.JSON(http.StatusOK, p)
}
