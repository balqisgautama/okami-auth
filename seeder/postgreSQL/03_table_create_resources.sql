-- +migrate Up
-- +migrate StatementBegin

-- USERS TABLE
CREATE SEQUENCE IF NOT EXISTS resource_id_pkey_sec
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE IF NOT EXISTS "resources" (
    "resource_id" BIGINT DEFAULT nextval('resource_id_pkey_sec'::regclass) NOT NULL PRIMARY KEY,
    "surname" VARCHAR(256) NOT NULL,
    "nickname" VARCHAR(256) NOT NULL UNIQUE,
    "client_id" VARCHAR(256) NOT NULL UNIQUE,
    "created_at" TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "created_client" VARCHAR(256) NULL,
    "updated_at" TIMESTAMP WITHOUT TIME ZONE,
    "updated_client" VARCHAR(256) NULL,
    "deleted_at" TIMESTAMP WITHOUT TIME ZONE,
    "deleted_client" VARCHAR(256) NULL
);

-- +migrate StatementEnd
-- +migrate Down
