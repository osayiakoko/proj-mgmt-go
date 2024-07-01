CREATE EXTENSION IF NOT EXISTS citext;

-- Create "users" table
CREATE TABLE "users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz(0) NOT NULL DEFAULT now(),
  "name" text NOT NULL,
  "email" citext NOT NULL,
  "password_hash" bytea NOT NULL,
  "activated" boolean NOT NULL,
  "version" integer NOT NULL DEFAULT 1,
  PRIMARY KEY ("id"),
  CONSTRAINT "users_email_key" UNIQUE ("email")
);
