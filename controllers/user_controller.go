package controllers

import (
	"auth-go/config"
	"auth-go/models"
	"auth-go/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserValidate struct {
	Username string `json:"username" validate:"required,min=2,max=20"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

func RegisterUser(c echo.Context) error {
	// request struct validation
	var user UserValidate

	// request params, and check body
	if err := c.Bind(&user); err != nil {
		utils.Logger.Error("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "invalid request body",
		})
	}

	// validation struc
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		utils.Logger.Error(errors)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": errors,
		})
	}

	// request struct model
	var userM models.User

	// check username is ready
	if err := config.DB.Where("username = ? ", user.Username).First(&userM).Error; err == nil {
		utils.Logger.Warn("Username " + user.Username + " is already")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":   400,
			"mesage": "Username " + user.Username + " is already",
		})
	}

	// validation password
	if err := utils.ValidatePassword(user.Password); err != nil {
		utils.Logger.Warn(err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": err.Error(),
		})
	}

	// hash password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.Logger.Error("Failed to hash password")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": "Failed to hash password",
		})
	}

	param := models.User{
		Username: user.Username,
		Password: hashedPassword,
	}

	// create to db
	if err := config.DB.Create(&param).Error; err != nil {
		utils.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"message": err.Error(),
		})
	}

	// return success
	utils.Logger.Info("Created successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "Created successfully",
		"data":    param,
	})
}
