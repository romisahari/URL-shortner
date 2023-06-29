package api

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"rom/database"

	"github.com/labstack/echo/v4"
)

func getShortURL(longURL string) string {
	h := sha256.New()
	h.Write([]byte(longURL))
	return hex.EncodeToString(h.Sum(nil))
}

func SaveURL(c echo.Context) error {
	longURL := c.Param("url")
	shortURL := getShortURL(longURL)[:16]
	database.PutURL(shortURL, longURL)
	return c.String(http.StatusOK, shortURL)
}

func GetURL(c echo.Context) error {
	str := database.GetURL(c.Param("url"))
	return c.Redirect(http.StatusTemporaryRedirect, str)
}
