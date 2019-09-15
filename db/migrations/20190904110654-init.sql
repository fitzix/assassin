-- +migrate Up
REVOKE ALL PRIVILEGES ON DATABASE "assassin" FROM PUBLIC;

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
    "id"       CHAR(16) PRIMARY KEY,
    "name"     VARCHAR(50) UNIQUE NOT NULL,
    "password" CHAR(32)           NOT NULL,
    "role_id"  INT                NOT NULL,
    "code"     INT DEFAULT 0      NOT NULL
);

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

DROP TABLE IF EXISTS "role";
CREATE TABLE "role" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) UNIQUE NOT NULL
);

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

DROP TABLE IF EXISTS "category";
CREATE TABLE "category" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) UNIQUE NOT NULL,
    "icon" VARCHAR(50)        NOT NULL
);

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

DROP TABLE IF EXISTS "tag";
CREATE TABLE "tag" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) UNIQUE NOT NULL
);

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

DROP TABLE IF EXISTS "app";
CREATE TABLE "app" (
    "id"         CHAR(16) PRIMARY KEY,
    "name"       VARCHAR(50)             NOT NULL,
    "type"       INT2      DEFAULT 0     NOT NULL,
    "icon"       VARCHAR(100)            NOT NULL,
    "title"      VARCHAR(200),
    "category"   INT       DEFAULT 1     NOT NULL,
    "created_at" TIMESTAMP               NOT NULL,
    "updated_at" TIMESTAMP               NOT NULL,
    "version_at" TIMESTAMP DEFAULT NOW() NOT NULL,
    "deleted_at" TIMESTAMP,
    "status"     BOOL      DEFAULT TRUE  NOT NULL
);
CREATE INDEX ON "app" ("type", "status", "deleted_at");

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

DROP TABLE IF EXISTS "app_carousel";
CREATE TABLE "app_carousel" (
    "id"     SERIAL PRIMARY KEY,
    "app_id" CHAR(16)    NOT NULL,
    "url"    VARCHAR(50) NOT NULL
);
CREATE INDEX ON "app_carousel" ("app_id");

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

DROP TABLE IF EXISTS "app_tag";
CREATE TABLE "app_tag" (
    "id"     SERIAL PRIMARY KEY,
    "app_id" CHAR(16) NOT NULL,
    "tag_id" INT      NOT NULL
);
CREATE INDEX ON "app_tag" ("app_id");

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

-- DROP TABLE IF EXISTS "app_category";
-- CREATE TABLE "app_category" (
--     "id"          SERIAL PRIMARY KEY,
--     "app_id"      CHAR(16) NOT NULL,
--     "category_id" INT      NOT NULL
-- );

-- CREATE INDEX ON "app_category" ("app_id");

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

DROP TABLE IF EXISTS "app_version";
CREATE TABLE "app_version" (
    "id"         SERIAL PRIMARY KEY,
    "app_id"     CHAR(16)          NOT NULL,
    "version"    VARCHAR(10)       NOT NULL,
    "size"       VARCHAR(10),
    "created_at" TIMESTAMP         NOT NULL,
    "status"     BOOL DEFAULT TRUE NOT NULL,
    CONSTRAINT " u_app_version " UNIQUE ("app_id", "version")
);
CREATE INDEX ON "app_version" ("app_id");

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

DROP TABLE IF EXISTS "download";
CREATE TABLE "download" (
    "id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL
);

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

DROP TABLE IF EXISTS "app_version_download";
CREATE TABLE "app_version_download" (
    "id"          SERIAL PRIMARY KEY,
    "app_version_id"  INT          NOT NULL,
    "download_id" INT          NOT NULL,
    "url"         VARCHAR(100) NOT NULL,
    "secret"      VARCHAR(50)
);
CREATE INDEX ON "app_version_download" ("app_version_id", "download_id");

-- -- -- -- -- -- -- -- -- -- -- -- -- -- --

DROP TABLE IF EXISTS "app_hot";
CREATE TABLE "app_hot" (
    "id"     SERIAL PRIMARY KEY,
    "app_id" CHAR(16)      NOT NULL,
    "hot"    INT DEFAULT 0 NOT NULL,
    "view"   INT DEFAULT 0 NOT NULL
);
CREATE INDEX ON "app_hot" ("app_id");

-- Column Comment

COMMENT ON COLUMN "app"."status" IS '0 正常 1 下架';
COMMENT ON COLUMN "app"."type" IS '0 应用 1 书籍';

COMMENT ON TABLE "app_hot" IS '下载和pv量统计表';
COMMENT ON COLUMN "app_hot"."hot" IS '下载量';
COMMENT ON COLUMN "app_hot"."view" IS '页面浏览量';


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