package controller

import (
	"comercial-api/db"
	"comercial-api/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Cliente struct {
	Nombre          string `json:"nombre"`
	ApellidoPaterno string `json:"apellido_materno"`
	ApellidoMaterno string `json:"apellido_paterno"`
	Ciudad          string `json:"ciudad"`
	Catergoria      int    `json:"categoria"`
}

func HandlerCreateClient(cfg db.ApiConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cliente Cliente
		err := c.Bind(&cliente)
		if err != nil {
			return c.String(400, "Error decoding request body")

		}
		fmt.Print(cliente)

		createClient, err := cfg.DB.CreateClient(context.Background(), database.CreateClientParams{
			ID:              uuid.New(),
			Nombre:          cliente.Nombre,
			ApellidoPaterno: cliente.ApellidoPaterno,
			ApellidoMaterno: cliente.ApellidoMaterno,
			Ciudad:          cliente.Ciudad,
			Categoria:       int32(cliente.Catergoria),
		})
		if err != nil {
			return c.String(500, err.Error())
		}
		return c.JSON(200, createClient)
	}

}
