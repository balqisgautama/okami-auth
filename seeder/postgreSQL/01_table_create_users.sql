-- +migrate Up
-- +migrate StatementBegin

-- USERS TABLE
CREATE SEQUENCE IF NOT EXISTS user_id_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS "users" (
    "user_id" BIGINT DEFAULT nextval('user_id_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "email" VARCHAR(256) NOT NULL UNIQUE,
    "password" VARCHAR(256) NOT NULL,
    "client_id" VARCHAR(256) NOT NULL UNIQUE,
    "status" SMALLINT NOT NULL DEFAULT 1,
    "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "created_client" VARCHAR(256) NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE,
    "updated_client" VARCHAR(256) NULL,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE,
    "deleted_client" VARCHAR(256) NULL
    );

-- +migrate StatementEnd
-- +migrate Down
