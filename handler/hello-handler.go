package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// HelloHandler - hello
func HelloHandler() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "はろー")
	}
}
