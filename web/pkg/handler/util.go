package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, component templ.Component) error {
	ctx := c.Request().Context()
	return component.Render(ctx, c.Response().Writer)
}

func redirect(c echo.Context, status int, url string) error {
	if len(c.Request().Header.Get("HX-Request")) > 0 {
		c.Response().Header().Set("HX-Redirect", url)
		return c.NoContent(status)
	}
	if status < 300 || status > 399 {
		return c.Redirect(http.StatusSeeOther, url)
	}
	return c.Redirect(status, url)
}
