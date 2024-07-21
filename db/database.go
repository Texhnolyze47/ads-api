package db

import "comercial-api/internal/database"

type Database interface {
	GetDb() *database.Queries
}
