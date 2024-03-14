-- name: CreateClient :one
INSERT INTO cliente (id, nombre, apellido_paterno, apellido_materno, ciudad, categoria)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;