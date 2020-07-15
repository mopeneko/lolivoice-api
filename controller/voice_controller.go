package controller

import (
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/lolivoice-api/interfaces"
	"github.com/mopeneko/lolivoice-api/model"
	"github.com/mopeneko/lolivoice-api/view"
	"net/http"
	"unicode/utf8"
)

func GetVoice(c echo.Context) error {
	req := new(interfaces.GetVoiceRequest)
	res := new(interfaces.GetVoiceResponse)

	if err := c.Bind(req); err != nil {
		return view.RenderGetVoice(c, http.StatusBadRequest, res)
	}

	if utf8.RuneCountInString(req.Text) >= 200 {
		return view.RenderGetVoice(c, http.StatusBadRequest, res)
	}

	bin, err := model.GenerateLoliVoice(req.Text)
	if err != nil {
		return view.RenderGetVoice(c, http.StatusInternalServerError, res)
	}

	encodedBin := base64.StdEncoding.EncodeToString(bin)
	res.Content = encodedBin
	return view.RenderGetVoice(c, http.StatusOK, res)
}

