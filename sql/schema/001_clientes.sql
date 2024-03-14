-- +goose Up
CREATE TABLE cliente (
    id               UUID PRIMARY KEY,
    nombre           VARCHAR(100) NOT NULL,
    apellido_paterno VARCHAR(100) NOT NULL,
    apellido_materno VARCHAR(100) NOT NULL,
    ciudad           VARCHAR(100) NOT NULL,
    categoria        INT          NOT NULL
);

-- +goose Down
DROP TABLE cliente;
