-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users
(
		id         SERIAL        NOT NULL,
    created_at BIGINT        NOT NULL,
    updated_at BIGINT        NOT NULL,
    deleted_at TIME,
    username   VARCHAR(255),
    email      VARCHAR(255)  NOT NULL,
    password   VARCHAR(255)  NOT NULL,
    firstname  VARCHAR(255),
    lastname   VARCHAR(255),
    CONSTRAINT users_pkey PRIMARY KEY (id)
    
    -- COMMENTED FOR SOFT DELETING
    -- CONSTRAINT users_email_unique UNIQUE (email),
    -- CONSTRAINT users_username_unique UNIQUE (username)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;