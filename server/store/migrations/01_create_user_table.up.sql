CREATE TABLE IF NOT EXISTS users
(
    id        SERIAL PRIMARY KEY,
    channel   VARCHAR(50)  NOT NULL UNIQUE,
    password  VARCHAR(64)  NOT NULL,
    email     VARCHAR(254) NOT NULL,
    join_date TIMESTAMP    NOT NULL DEFAULT NOW()
);
