create table if not exists USERS
(
    ID        serial primary key,
    CHANNEL   varchar(50)  not null unique,
    PASSWORD  varchar(64)  not null,
    EMAIL     varchar(254) not null,
    JOIN_DATE timestamp    not null default now()
);
