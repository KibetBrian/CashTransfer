CREATE TYPE "Currency" AS ENUM (
  'USD'
);

CREATE TABLE "users" (
  "id" uuid UNIQUE NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone" varchar,
  "createdAt" timestamptz NOT NULL DEFAULT (now()),
  "updatedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" uuid PRIMARY KEY NOT NULL,
  "holderId" uuid NOT NULL,
  "balance" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "Id" uuid NOT NULL,
  "sender" uuid NOT NULL,
  "reciever" uuid NOT NULL,
  "amount" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("holderId") REFERENCES "users" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("sender") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("reciever") REFERENCES "accounts" ("id");

CREATE INDEX ON "users" ("id");

CREATE UNIQUE INDEX ON "accounts" ("holderId");

CREATE INDEX ON "transactions" ("sender");

CREATE INDEX ON "transactions" ("reciever");

CREATE INDEX ON "transactions" ("Id");
