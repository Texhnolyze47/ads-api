-- name: CreateCommercial :one
INSERT INTO comercial (id, nombre, apellido_paterno, apellido_materno,ciudad, commision)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;