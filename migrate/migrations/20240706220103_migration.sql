-- Create "permissions" table
CREATE TABLE "permissions" (
  "id" bigserial NOT NULL,
  "code" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create "users_permissions" table
CREATE TABLE "users_permissions" (
  "user_id" bigint NOT NULL,
  "permission_id" bigint NOT NULL,
  PRIMARY KEY ("user_id", "permission_id"),
  CONSTRAINT "users_permissions_permission_id_fkey" FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id") ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT "users_permissions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE CASCADE
);

-- Add the two permissions to the table.
INSERT INTO permissions (code) 
VALUES
  ('tasks:read'),
  ('tasks:write');
