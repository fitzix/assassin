-- +migrate Up

DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
    "id"       SERIAL PRIMARY KEY,
    "uid"      CHAR(16) UNIQUE    NOT NULL,
    "name"     VARCHAR(50) UNIQUE NOT NULL,
    "password" CHAR(60)           NOT NULL,
    "role_id"  INT                NOT NULL,
    "code"     INT  DEFAULT 0     NOT NULL,
    "status"   BOOL DEFAULT TRUE  NOT NULL
);


DROP TABLE IF EXISTS "role";
CREATE TABLE "role" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) UNIQUE NOT NULL
);


DROP TABLE IF EXISTS "category";
CREATE TABLE "category" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) UNIQUE NOT NULL,
    "icon" VARCHAR(50)        NOT NULL
);

DROP TABLE IF EXISTS "tag";
CREATE TABLE "tag" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) UNIQUE NOT NULL
);


DROP TABLE IF EXISTS "app";
CREATE TABLE "app" (
    "id"         SERIAL PRIMARY KEY,
    "app_id"     CHAR(16) UNIQUE         NOT NULL,
    "name"       VARCHAR(50)             NOT NULL,
    "type"       INT2      DEFAULT 0     NOT NULL,
    "icon"       VARCHAR(100)            NOT NULL,
    "title"      VARCHAR(200),
    "category"   INT2      DEFAULT 1     NOT NULL,
    "created_at" TIMESTAMP               NOT NULL,
    "updated_at" TIMESTAMP               NOT NULL,
    "version_at" TIMESTAMP DEFAULT NOW() NOT NULL,
    "deleted_at" TIMESTAMP,
    "status"     BOOL      DEFAULT TRUE  NOT NULL
);
CREATE INDEX ON "app" ("type", "status", "deleted_at");


DROP TABLE IF EXISTS "app_tag";
CREATE TABLE "app_tag" (
    "id"     SERIAL PRIMARY KEY,
    "app_id" CHAR(16) NOT NULL,
    "tag_id" INT      NOT NULL
);
CREATE INDEX ON "app_tag" ("app_id", "tag_id");


DROP TABLE IF EXISTS "carousel";
CREATE TABLE "carousel" (
    "id"     SERIAL PRIMARY KEY,
    "app_id" CHAR(16)    NOT NULL,
    "url"    VARCHAR(50) NOT NULL
);
CREATE INDEX ON "carousel" ("app_id");


DROP TABLE IF EXISTS "app_category";
CREATE TABLE "app_category" (
    "id"          SERIAL PRIMARY KEY,
    "app_id"      CHAR(16) NOT NULL,
    "category_id" INT      NOT NULL
);

CREATE INDEX ON "app_category" ("app_id", "category_id");


DROP TABLE IF EXISTS "version";
CREATE TABLE "version" (
    "id"         SERIAL PRIMARY KEY,
    "app_id"     CHAR(16)          NOT NULL,
    "name"       VARCHAR(10)       NOT NULL,
    "version"    VARCHAR(10)       NOT NULL,
    "size"       INT8 DEFAULT 0    NOT NULL,
    "created_at" TIMESTAMP         NOT NULL,
    "status"     BOOL DEFAULT TRUE NOT NULL,
    CONSTRAINT "unique_app_version" UNIQUE ("app_id", "version")
);
CREATE INDEX ON "version" ("app_id", "version");


DROP TABLE IF EXISTS "provider";
CREATE TABLE "provider" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL
);


DROP TABLE IF EXISTS "source";
CREATE TABLE "source" (
    "id"          SERIAL PRIMARY KEY,
    "version_id"  INT          NOT NULL,
    "provider_id" INT          NOT NULL,
    "url"         VARCHAR(100) NOT NULL,
    "secret"      VARCHAR(50)
);
CREATE INDEX ON "source" ("version_id", "provider_id");


DROP TABLE IF EXISTS "hot";
CREATE TABLE "hot" (
    "id"     SERIAL PRIMARY KEY,
    "app_id" CHAR(16)      NOT NULL,
    "hot"    INT DEFAULT 0 NOT NULL,
    "view"   INT DEFAULT 0 NOT NULL
);
CREATE INDEX ON "hot" ("app_id");
COMMENT ON TABLE "hot" IS '下载和pv量统计表';
COMMENT ON COLUMN "hot"."hot" IS '下载量';
COMMENT ON COLUMN "hot"."view" IS '页面浏览量';

-- init data
INSERT INTO "role" ("name")
VALUES ('默认角色');
INSERT INTO "user" ("uid", "name", "password", "role_id")
VALUES ('5-sT6lAKM3_weQIE', 'admin', '$2a$10$kX4Ianl/Ua4CLjxHLDi84OxsN6dN234DAwB2JRyJAyUP8mmz.cc2S', 1);


-- +migrate Down
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS "role";
DROP TABLE IF EXISTS "category";
DROP TABLE IF EXISTS "tag";
DROP TABLE IF EXISTS "app";
DROP TABLE IF EXISTS "app_tag";
DROP TABLE IF EXISTS "provider";
DROP TABLE IF EXISTS "app_category";
DROP TABLE IF EXISTS "version";
DROP TABLE IF EXISTS "source";
DROP TABLE IF EXISTS "hot";
DROP TABLE IF EXISTS "carousel";
