-- +goose Up
CREATE TABLE pedido (
    id UUID PRIMARY KEY,
    cantidad DOUBLE PRECISION NOT NULL,
    fecha TIMESTAMP NOT NULL,
    id_cliente UUID NOT NULL REFERENCES cliente (id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE pedido;