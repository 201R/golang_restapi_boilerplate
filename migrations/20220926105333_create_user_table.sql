-- create "users" table
CREATE TABLE "users" ("id" character varying NOT NULL, "first_name" character varying NULL, "last_name" character varying NULL, "email" character varying NOT NULL, "location" character varying NULL, "phone" character varying NULL, "password" character varying NOT NULL, "status" boolean NOT NULL DEFAULT true, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
-- create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX "users_email_key" ON "users" ("email");
