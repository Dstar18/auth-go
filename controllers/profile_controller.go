package controllers

import (
	"auth-go/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProfileBasicAuth(c echo.Context) error {
	utils.Logger.Info("Basic Auth Successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Basic Auth Successfully",
	})
}

func ProfileJWT(c echo.Context) error {
	utils.Logger.Info("JWT Successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "JWT Successfully",
	})
}

func Dashboard(c echo.Context) error {
	username := c.Get("username").(string)
	utils.Logger.Info("Anda berhasil login menggunakan user " + username)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Auth Session Successfully",
		"data":    "Anda berhasil login menggunakan user " + username,
	})
}
