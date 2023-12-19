Creación de la DB
```
CREATE TABLE clientes (
    id INT NOT NULL AUTO_INCREMENT,
    nombre VARCHAR(100) NOT NULL,
    correo VARCHAR(100) NOT NULL,
    telefono VARCHAR(20) NOT NULL,
    fecha DATETIME NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```


Insertar algún cliente

```
INSERT INTO clientes (nombre, correo, telefono, fecha) VALUES
('Juan Pérez', 'juan.perez@example.com', '555-1234', '2023-12-16 12:00:00'),
('María González', 'maria.gonzalez@example.com', '555-5678', '2023-12-16 12:00:00'),
('Carlos Rodríguez', 'carlos.rodriguez@example.com', '555-9101', '2023-12-16 12:00:00'),
('Ana García', 'ana.garcia@example.com', '555-1122', '2023-12-16 12:00:00'),
('Eduardo Vargas', 'eduardo.vargas@example.com', '555-3344', '2023-12-16 12:00:00'),
('Laura Martínez', 'laura.martinez@example.com', '555-5566', '2023-12-16 12:00:00'),
('Alejandro López', 'alejandro.lopez@example.com', '555-7788', '2023-12-16 12:00:00'),
('Isabel Torres', 'isabel.torres@example.com', '555-9900', '2023-12-16 12:00:00'),
('Miguel Sánchez', 'miguel.sanchez@example.com', '555-1122', '2023-12-16 12:00:00'),
('Luisa Morales', 'luisa.morales@example.com', '555-3344', '2023-12-16 12:00:00'),
('Roberto Hernández', 'roberto.hernandez@example.com', '555-5566', '2023-12-16 12:00:00'),
('Sofía Jiménez', 'sofia.jimenez@example.com', '555-7788', '2023-12-16 12:00:00'),
('Pedro Castro', 'pedro.castro@example.com', '555-9900', '2023-12-16 12:00:00'),
('Verónica Ramírez', 'veronica.ramirez@example.com', '555-1122', '2023-12-16 12:00:00'),
('Javier Ruiz', 'javier.ruiz@example.com', '555-3344', '2023-12-16 12:00:00'),
('Carmen Flores', 'carmen.flores@example.com', '555-5566', '2023-12-16 12:00:00'),
('Raúl Díaz', 'raul.diaz@example.com', '555-7788', '2023-12-16 12:00:00'),
('Lorena Gómez', 'lorena.gomez@example.com', '555-9900', '2023-12-16 12:00:00'),
('Gabriel Nuñez', 'gabriel.nunez@example.com', '555-1122', '2023-12-16 12:00:00'),
('Adriana Ortega', 'adriana.ortega@example.com', '555-3344', '2023-12-16 12:00:00');

```



Instalar godotenv


```
go get github.com/joho/godotenv
```

Instalar mysql

```
go get github.com/go-sql-driver/mysql
```