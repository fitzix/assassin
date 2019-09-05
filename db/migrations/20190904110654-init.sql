-- +migrate Up
REVOKE ALL PRIVILEGES ON DATABASE "assassin" FROM PUBLIC;

DROP TABLE IF EXISTS "user";

CREATE TABLE "user" (
    "id"       CHAR(16) PRIMARY KEY,
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
    "id"          CHAR(16) PRIMARY KEY,
    "name"        VARCHAR(50) UNIQUE NOT NULL,
    "description" TEXT,
    "created_at"  TIMESTAMP          NOT NULL,
    "updated_at"  TIMESTAMP          NOT NULL,
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
    "app_id"      CHAR(16) NOT NULL,
    "category_id" INT      NOT NULL
);

DROP TABLE IF EXISTS "app_version";

CREATE TABLE "app_version" (
    "id"          SERIAL PRIMARY KEY,
    "app_id"      CHAR(16)     NOT NULL,
    "version"     VARCHAR(10)  NOT NULL,
    "size"        VARCHAR(10),
    "download_id" INT2         NOT NULL,
    "url"         VARCHAR(100) NOT NULL,
    "secret"      VARCHAR(50),
    "created_at"  TIMESTAMP    NOT NULL,
    CONSTRAINT "u_app_version" UNIQUE ("app_id", "version")
);

DROP TABLE IF EXISTS "download";

CREATE TABLE "download" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL
);

DROP TABLE IF EXISTS "app_hot";

CREATE TABLE "app_hot" (
    "id"     SERIAL PRIMARY KEY,
    "app_id" CHAR(16)      NOT NULL,
    "hot"    INT DEFAULT 0 NOT NULL,
    "view"   INT DEFAULT 0 NOT NULL
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
DROP TABLE IF EXISTS "download";
DROP TABLE IF EXISTS "app_hot";