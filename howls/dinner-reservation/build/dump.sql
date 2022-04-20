CREATE DATABASE IF NOT EXISTS `reservation`;
USE `reservation`;
CREATE TABLE IF NOT EXISTS tables
(
    id    INT AUTO_INCREMENT
        PRIMARY KEY,
    seats INT NOT NULL,
    CONSTRAINT table_id_uindex
        UNIQUE (id)
);

CREATE TABLE IF NOT EXISTS clients
(
    id        INT AUTO_INCREMENT
        PRIMARY KEY,
    name      VARCHAR(50) NOT NULL,
    groupSize INT         NOT NULL,
    checkIn   DATETIME    NULL,
    tableId   INT         NULL,
    CONSTRAINT clients_name_uindex
        UNIQUE (name),
    CONSTRAINT clients_tables_id_fk
        FOREIGN KEY (tableId) REFERENCES tables (id)
            ON DELETE CASCADE
);


