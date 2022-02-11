CREATE TYPE "Currency" AS ENUM (
  'USD'
);

CREATE TABLE "users" (
  "id" uuid NOT NULL,
  "name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "phone" varchar,
  "createdAt" timestamptz DEFAULT (now()),
  "updatedAt" timestamptz DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" uuid PRIMARY KEY NOT NULL,
  "holderId" uuid NOT NULL,
  "balance" decimal NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Transaction" (
  "transactionId" uuid NOT NULL,
  "from" uuid NOT NULL,
  "to" uuid NOT NULL,
  "amount" decimal NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("holderId") REFERENCES "users" ("id");

ALTER TABLE "Transaction" ADD FOREIGN KEY ("from") REFERENCES "accounts" ("id");

ALTER TABLE "Transaction" ADD FOREIGN KEY ("to") REFERENCES "accounts" ("id");

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "accounts" ("holderId");

CREATE INDEX ON "Transaction" ("from");

CREATE INDEX ON "Transaction" ("to");

CREATE INDEX ON "Transaction" ("transactionId");
