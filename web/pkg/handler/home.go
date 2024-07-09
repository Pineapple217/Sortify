package handler

import (
	"github.com/Pineapple217/Sortify/web/pkg/auth"
	"github.com/Pineapple217/Sortify/web/pkg/view"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Home(c echo.Context) error {
	a := auth.GetAuth(c.Request().Context())
	if !a.CheckSpotify() {
		return render(c, view.Home("a"))
	}
	client := a.GetClient(c.Request().Context(), h.SpotifyAuth, h.DB)
	user, err := client.CurrentUser(c.Request().Context())
	if err != nil {
		return err
	}
	return render(c, view.Home(user.DisplayName))
}
