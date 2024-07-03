-- Create "tokens" table
CREATE TABLE "tokens" (
  "user_id" bigint NOT NULL,
  "expiry" timestamptz NOT NULL,
  "hash" bytea NOT NULL,
  "scope" text NOT NULL,
  PRIMARY KEY ("hash"),
  CONSTRAINT "tokens_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);
