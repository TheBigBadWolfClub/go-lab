CREATE DATABASE cleanDB;
USE cleanDB;

create table if not exists customers
(
    id       int auto_increment,
    name     varchar(20) not null,
    contract int         not null,
    constraint customers_id_uindex
        unique (id)
);

ALTER TABLE customers
    AUTO_INCREMENT = 10000;
alter table customers
    add primary key (id);

insert into cleanDB.customers (id, name, contract)
values (10001, 'Layer&Layer', 1),
       (10002, 'Repair Office', 0),
       (10003, 'Cleaning Services', 2),
       (10004, 'International School', 0);

create table if not exists contracts
(
    id        int         not null,
    type      varchar(20) not null,
    max       int         not null,
    promotion int         not null,
    constraint contracts_type_uindex
        unique (type)
);
insert into cleanDB.contracts (id, type, max, promotion)
values (1, 'basic', 10, 5),
       (2, 'standard', 25, 7),
       (3, 'premium', 100, 18),
       (4, 'trial', 5, 50);


create table if not exists power_tools
(
    code       varchar(10) not null,
    type       varchar(50) not null,
    rate       int         not null,
    assignment int default null,
    constraint power_tools_code_uindex
        unique (code),
    constraint power_tools_customers_id_fk
        foreign key (assignment) references customers (id)
            on delete set null
);

alter table power_tools
    add primary key (code);


insert into cleanDB.power_tools (code, type, rate)
values ('12-clean-a', 'drill', 11),
       ('12-clean-e', 'drill', 10),
       ('4-sander-a', 'sander', 10),
       ('4-sander-b', 'sander', 32),
       ('5346-a3', 'grinder', 2),
       ('5346-a4', 'grinder', 2),
       ('56-tg-e', 'saw', 5),
       ('56-tg-g', 'saw', 5);

create table if not exists assign_log
(
    tool       varchar(10)          not null,
    customer   int                  not null,
    start      datetime,
    end        datetime,
    liquidated tinyint(1) default 0 null,
    constraint assign_log_customers_id_fk
        foreign key (customer) references customers (id),
    constraint assign_log_power_tools_code_fk
        foreign key (tool) references power_tools (code)
);