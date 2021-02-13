package main

import (
	"fmt"
	"github.com/auth/backend/handlers"
	middle "github.com/auth/backend/middleware"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	handler, err := handlers.NewHTTPHandler()
	if err != nil {
		log.Fatal(err)
	}
	server := echo.New()
	server.Use(
		middle.Cors(),
	)

	// index route
	server.GET("/", handler.Index)

	// auth routes
	server.GET("/auth/providers", handler.Providers)
	server.POST("/auth/login", handler.Login)

	// user management routes
	server.GET("/user", handler.GetUser)

	// test private (authenticated) and public routes
	server.GET("/private", handler.Private)
	server.GET("/public", handler.Public)

	log.Println("listening on localhost:4061")
	bind := fmt.Sprintf(":4061")
	log.Fatal(server.Start(bind))
}


