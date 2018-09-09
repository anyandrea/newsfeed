-- user
CREATE TABLE IF NOT EXISTS user (
    pk_user_id        INTEGER NOT NULL AUTO_INCREMENT,
    password          VARCHAR(255) NOT NULL,
    name              VARCHAR(255) NOT NULL UNIQUE,
    email             VARCHAR(255) NOT NULL UNIQUE,
    role              VARCHAR(10) NOT NULL,
    active            BOOLEAN,
    PRIMARY KEY (pk_user_id)
);
