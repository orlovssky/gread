-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE links
(
		id           SERIAL        NOT NULL,
    created_at   BIGINT        NOT NULL,
    updated_at   BIGINT        NOT NULL,
    deleted      SMALLINT      NOT NULL,
    user_id      INT           NOT NULL,
    title        VARCHAR(255),
    author       VARCHAR(255),
    content      TEXT,
    text_content TEXT,
    length       BIGINT,
    excerpt      TEXT,
    site_name    VARCHAR(255),
    image        TEXT,
    favicon      TEXT,
    url          TEXT,
    CONSTRAINT links_pkey PRIMARY KEY (id)
    
    -- COMMENTED FOR SOFT DELETING
    -- CONSTRAINT users_email_unique UNIQUE (email),
    -- CONSTRAINT users_username_unique UNIQUE (username)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE links;