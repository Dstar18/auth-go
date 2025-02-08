package middleware

import (
	"auth-go/config"
	"auth-go/models"
	"auth-go/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func BasicAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			// check authentication
			username, password, ok := c.Request().BasicAuth()
			if !ok {
				utils.Logger.Error("Unauthorized")
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    401,
					"message": "Unauthorized",
				})
			}

			// check user on database
			var user models.User
			if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
				utils.Logger.Warn("Invalid username or password")
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    401,
					"message": "Invalid username or password",
				})
			}

			// check password in hash
			if err := utils.CheckPassword(user.Password, password); err != nil {
				utils.Logger.Warn(err.Error())
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    401,
					"message": err.Error(),
				})
			}

			return next(c)
		}
	}

}
