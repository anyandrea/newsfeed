-- feed
CREATE TABLE IF NOT EXISTS feed (
    pk_feed_id      INTEGER NOT NULL AUTO_INCREMENT,
    title           VARCHAR(64) NOT NULL,
    link            VARCHAR(128) NOT NULL UNIQUE,
    updated         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    fetched         TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (pk_feed_id)
);

-- feed item
CREATE TABLE IF NOT EXISTS item (
    pk_feed_item_id     INTEGER NOT NULL AUTO_INCREMENT,
    fk_feed_id          INTEGER NOT NULL,
    title               VARCHAR(64) NOT NULL,
    link                VARCHAR(128) NOT NULL,
    updated             TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    published           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (pk_feed_item_id),
    FOREIGN KEY (fk_feed_id) REFERENCES feed (pk_feed_id) ON DELETE CASCADE
);

-- subscription
CREATE TABLE IF NOT EXISTS subscription (
    fk_feed_id      INTEGER NOT NULL,
    fk_user_id      INTEGER NOT NULL,
    show_entries    INTEGER NOT NULL DEFAULT 10,
    PRIMARY KEY (fk_feed_id, fk_user_id),
    FOREIGN KEY (fk_feed_id) REFERENCES feed (pk_feed_id) ON DELETE CASCADE,
    FOREIGN KEY (fk_user_id) REFERENCES user (pk_user_id) ON DELETE CASCADE
);
