CREATE TABLE IF NOT EXISTS "users" (
    "id" SERIAL PRIMARY KEY,
    "userName" VARCHAR (64) NOT NULL,
    "firstName" VARCHAR (64) NOT NULL,
    "lastName" VARCHAR (64) NOT NULL DEFAULT (NOW()),
    "created_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
    "updated_at" TIMESTAMP NOT NULL DEFAULT (NOW()),
)