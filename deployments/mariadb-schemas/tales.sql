CREATE DATABASE tales;
USE tales;

create table if not exists Tales
(
    id       int auto_increment,
    name     varchar(20) not null,
    extid    int,
    constraint customers_id_uindex
        unique (id)
);

ALTER TABLE Tales
    AUTO_INCREMENT = 10000;
alter table Tales
    add primary key (id);

insert into tales.Tales ( name, extid)
values ('Little Red Riding Hood', 1);