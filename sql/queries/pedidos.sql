-- name: ListProducts :many
SELECT * FROM pedido
ORDER BY fecha DESC;

-- name: ListProductsLimitTwo :many
SELECT * FROM pedido ORDER BY cantidad DESC LIMIT 2;
-- name: ListIDClientMakePurchase :many
SELECT DISTINCT id_cliente FROM pedido;

-- name: ListProducts2007 :many
SELECT * FROM pedido WHERE  EXTRACT(YEAR FROM fecha) = 2017 AND cantidad > 500;

-- name: AdCommission :many
SELECT nombre, apellido_paterno, apellido_materno FROM comercial WHERE commision BETWEEN 0.05 AND 0.11;


-- name: MaxCommission :one
SELECT MAX(commision) FROM comercial;


-- name: LastNameNotNull :many
SELECT id, nombre, apellido_paterno FROM cliente
WHERE apellido_materno IS NOT NULL AND LENGTH(apellido_materno) > 0
ORDER BY apellido_paterno, apellido_materno, nombre;

-- name: AStartName :many
SELECT nombre FROM cliente WHERE (nombre LIKE 'A%' AND nombre LIKE '%n') OR nombre LIKE 'P%';

-- name: ANoStartName :many
SELECT nombre FROM cliente WHERE nombre NOT LIKE 'A%';

-- name: ONameADName :many
SELECT DISTINCT nombre FROM comercial WHERE nombre LIKE '%o';

-- name: CreateOrder :one
INSERT INTO pedido (id,cantidad,fecha,id_cliente,id_comercial)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;
