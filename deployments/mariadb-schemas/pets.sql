CREATE DATABASE creatures;
USE creatures;

create table if not exists Pets
(
    id       int auto_increment,
    name     varchar(20) not null,
    type     varchar(20) not null,
    constraint customers_id_uindex
        unique (id)
);

ALTER TABLE Pets
    AUTO_INCREMENT = 10000;
alter table Pets
    add primary key (id);

insert into creatures.Pets ( name, type)
values ('Wag','Dog'),
       ('Meow','Cat'),
       ('Bark','Dog');