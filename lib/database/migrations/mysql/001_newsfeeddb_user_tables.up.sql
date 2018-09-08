-- user
CREATE TABLE IF NOT EXISTS user (
    pk_user_id        INTEGER NOT NULL AUTO_INCREMENT,
    password          VARCHAR(64) NOT NULL,
    name              VARCHAR(32) NOT NULL UNIQUE,
    email             VARCHAR(64) NOT NULL UNIQUE,
    role              VARCHAR(10) NOT NULL,
    PRIMARY KEY (pk_user_id)
);
