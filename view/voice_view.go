package view

import (
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/lolivoice-api/interfaces"
)

func RenderGetVoice(c echo.Context, code int, res *interfaces.GetVoiceResponse) error {
	return c.JSON(code, res)
}

