-- -- Create "users" table
-- CREATE TABLE "public"."users" (
--   "id" bigint NOT NULL, 
--   "name" character varying NULL, 
--   PRIMARY KEY ("id")
-- );

-- -- create "user_profile" table
-- CREATE TABLE "public"."user_profiles" (
--   "id" bigint NOT NULL,
--   "bio" text NULL,
--   "user_id" bigint NULL,
--   PRIMARY KEY ("id"),
--   CONSTRAINT "user_fk" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id")
-- );

-- -- create "blog_posts" table
-- CREATE TABLE "public"."blog_posts" (
--   "id" bigint NOT NULL,
--   "title" character varying(100) NULL,
--   "body" text NULL,
--   "author_id" bigint NULL,
--   PRIMARY KEY ("id"),
--   CONSTRAINT "author_fk" FOREIGN KEY ("author_id") REFERENCES "public"."users" ("id")
-- );


-- create "blog_posts" table
CREATE TABLE tasks (
  id bigserial PRIMARY KEY,
  title text NOT NULL,
  description text NOT NULL,
  priority text NOT NULL,
  status text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT NOW(),
  updated_at timestamptz NOT NULL DEFAULT NOW()
  -- created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
);

CREATE INDEX IF NOT EXISTS tasks_title_idx ON tasks USING GIN (to_tsvector('simple', title)); 
CREATE INDEX IF NOT EXISTS tasks_priority_idx ON tasks USING HASH (priority);
CREATE INDEX IF NOT EXISTS tasks_status_idx ON tasks USING HASH (status);

CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(), name text NOT NULL,
  email citext UNIQUE NOT NULL,
  password_hash bytea NOT NULL,
  activated bool NOT NULL,
  version integer NOT NULL DEFAULT 1
);


CREATE TABLE IF NOT EXISTS tokens (
  user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE,
  expiry timestamptz NOT NULL,
  hash bytea PRIMARY KEY,
  scope text NOT NULL
);

CREATE TABLE IF NOT EXISTS permissions ( 
  id bigserial PRIMARY KEY,
  code text NOT NULL
);

CREATE TABLE IF NOT EXISTS users_permissions (
  user_id bigint NOT NULL REFERENCES users ON DELETE CASCADE, 
  permission_id bigint NOT NULL REFERENCES permissions ON DELETE CASCADE, 
  PRIMARY KEY (user_id, permission_id)
);
