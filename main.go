package main

import (
	"comercial-api/db"
	"comercial-api/internal/database"
	"comercial-api/router"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func main() {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	godotenv.Load(".env")

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		e.Logger.Fatal("DB_URL not found")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		e.Logger.Fatal("Error connecting to database")
	}

	apiCfg := db.Database{
		DB: database.New(conn),
	}

	r, err := router.SetupRouter(apiCfg, e)

	if err != nil {
		e.Logger.Fatal("Error setting up router")
	} else {
		e = r
	}
}
