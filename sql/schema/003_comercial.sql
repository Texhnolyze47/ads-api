-- +goose Up
CREATE TABLE comercial (
    id UUID PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    apellido_paterno VARCHAR(100) NOT NULL,
    apellido_materno VARCHAR(100) NOT NULL,
    ciudad           VARCHAR(100) NOT NULL,
    commision FLOAT NOT NULL

);

ALTER TABLE pedido ADD COLUMN id_comercial UUID NOT NULL REFERENCES comercial (id) ON DELETE CASCADE;
-- +goose Down
ALTER TABLE pedido DROP COLUMN id_comercial;
DROP TABLE comercial;
