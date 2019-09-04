-- +migrate Up
REVOKE ALL PRIVILEGES ON DATABASE "assassin" FROM PUBLIC;

DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
    "id"       SERIAL PRIMARY KEY,
    "name"     VARCHAR(50) UNIQUE  NOT NULL,
    "password" VARCHAR(50)         NOT NULL,
    "email"    VARCHAR(300) UNIQUE NOT NULL,
    "role_id"  INT                 NOT NULL
);

DROP TABLE IF EXISTS "role";

CREATE TABLE "role" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) UNIQUE NOT NULL
);

DROP TABLE IF EXISTS "category";

CREATE TABLE "category" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) UNIQUE NOT NULL
);

DROP TABLE IF EXISTS "tag";

CREATE TABLE "tag" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) UNIQUE NOT NULL
);

DROP TABLE IF EXISTS "app";

CREATE TABLE "app" (
    "id"          SERIAL PRIMARY KEY,
    "name"        VARCHAR(50) UNIQUE NOT NULL,
    "description" TEXT,
    "created_at"  TIMESTAMP          NOT NULL,
    "updated_at"  TIMESTAMP,
    "deleted_at"  TIMESTAMP
);

DROP TABLE IF EXISTS "app_tag";

CREATE TABLE "app_tag" (
    "id"     SERIAL PRIMARY KEY,
    "app_id" INT NOT NULL,
    "tag_id" INT NOT NULL
);

DROP TABLE IF EXISTS "app_category";

CREATE TABLE "app_category" (
    "id"          SERIAL PRIMARY KEY,
    "app_id"      INT NOT NULL,
    "category_id" INT NOT NULL
);

DROP TABLE IF EXISTS "app_version";

CREATE TABLE "app_version" (
    "id"      SERIAL PRIMARY KEY,
    "app_id"  INT         NOT NULL,
    "version" VARCHAR(10) NOT NULL,
    CONSTRAINT "u_app_version" UNIQUE ("app_id", "version")
);
-- +migrate Down
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "role";
DROP TABLE IF EXISTS "category";
DROP TABLE IF EXISTS "tag";
DROP TABLE IF EXISTS "app";
DROP TABLE IF EXISTS "app_tag";
DROP TABLE IF EXISTS "app_version";
DROP TABLE IF EXISTS "app_category";