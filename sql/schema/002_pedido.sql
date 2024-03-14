-- +goose Up
CREATE TABLE pedido (
    id INT PRIMARY KEY,
    cantidad DOUBLE PRECISION NOT NULL,
    fecha DATE,
    id_cliente INT CONSTRAINT fk_cliente FOREIGN KEY (id_cliente)  REFERENCES cliente (id)

)