package controller

import "github.com/labstack/echo/v4"

func AddCliente(echo.Context) error {

	type Cliente struct {
		Id              int    `json:"id"`
		Nombre          string `json:"nombre"`
		ApellidoPaterno string `json:"apellido-materno"`
		ApellidoMaterno string `json:"apellido-paterno"`
		Ciudad          string `json:"ciudad"`
		Catergoria      string `json:"categoria"`
	}

	return nil
}
