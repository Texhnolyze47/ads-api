package controller

import (
	"comercial-api/internal/db"
)

type apiConfig struct {
	DB *db.Queries
}

func (apiCfg *apiConfig) CreateClient() error {

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
