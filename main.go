package main

import (
	"auth-go/config"
	"auth-go/controllers"
	"auth-go/middleware"
	"auth-go/utils"

	"github.com/labstack/echo/v4"
)

func main() {

	// Initialize logger
	utils.InitLogger()

	// connection to db
	config.InitDB()
	// config.DB.AutoMigrate(&models.User{})

	// initialize echo
	e := echo.New()

	// Middleware Basic Auth
	basicAuth := middleware.BasicAuthMiddleware()

	protectedBasicAuth := e.Group("/basic")
	protectedBasicAuth.Use(basicAuth)

	// Middleware JWT
	protectedJWT := e.Group("/jwt")
	protectedJWT.Use(middleware.JWTMiddleware)

	// Public Route
	e.POST("/register", controllers.RegisterUser)
	e.POST("/login", controllers.LoginUser)

	// Protected Route
	protectedBasicAuth.GET("/info", controllers.ProfileBasicAuth)
	protectedJWT.GET("/notif", controllers.ProfileJWT)

	// start server with logging
	e.Logger.Fatal(e.Start(":3000"))
}
