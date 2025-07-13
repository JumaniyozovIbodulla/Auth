CREATE TYPE "statuses" AS ENUM('active','deleted');
CREATE TYPE "otp_statuses" AS ENUM('unconfirmed','confirmed');


CREATE TABLE IF NOT EXISTS "users" (
    "id" UUID PRIMARY KEY,
    "status"     statuses, -- [active, deleted]
    "name"       VARCHAR(255) NOT NULL,
    "email"      VARCHAR(255) NOT NULL,
    "password"   TEXT NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT NOW(),
    "created_by" UUID
);

CREATE TABLE IF NOT EXISTS "otp" (
    "id"         UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "email"      VARCHAR(255) NOT NULL,
    "status"     otp_statuses, -- [unconfirmed, confirmed]
    "code"       VARCHAR(50)  NOT NULL,
    "expires_at" TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS "sysusers" (
    "id"         UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "status"     statuses  NOT NULL, -- [active, deleted]
    "name"       VARCHAR(255) NOT NULL,
    "password"   VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT NOW(),
    "created_by" UUID
);

CREATE TABLE IF NOT EXISTS "sysuser_roles" (
    "id"         UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "sysuser_id" UUID NOT NULL,
    "role_id"    UUID NOT NULL
);

CREATE TABLE IF NOT EXISTS "roles" (
    "id"         UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    "status"     statuses  NOT NULL, -- [active, deleted]
    "name"       VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMPTZ DEFAULT NOW(),
    "created_by" UUID
);