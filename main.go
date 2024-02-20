package main

import (
	"comercial-api/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {

	e := echo.New()

	godotenv.Load(".env")

	port := os.Getenv("DB_URL")

	if port == "" {
		port = "3000"
	}

	r, err := router.SetupRouter(e)

	if err != nil {
		e.Logger.Fatal("Error setting up router")
	} else {
		e = r
	}
}
