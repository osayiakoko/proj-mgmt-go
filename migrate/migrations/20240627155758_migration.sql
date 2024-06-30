-- Create "todos" table
CREATE TABLE "todos" (
  "id" bigserial NOT NULL,
  "title" text NOT NULL,
  "description" text NOT NULL,
  "priority" text NOT NULL,
  "status" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);
