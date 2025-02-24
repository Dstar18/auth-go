package middleware

import (
	"auth-go/config"
	"auth-go/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SessionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		// pengecekkan session dan username
		session, _ := config.Store.Get(c.Request(), "session")
		username, ok := session.Values["username"].(string)

		// jika tidak ditemukan, maka dikembalikan pesan error
		if !ok || username == "" {
			utils.Logger.Error("Unauthorized")
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": "Unauthorized",
			})
		}

		// jika berhasil, maka diteruskan dengan menyimpan data username yang dapat dipanggil dari mana saja
		c.Set("username", username)
		return next(c)
	}
}
