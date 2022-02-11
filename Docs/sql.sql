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

CREATE TABLE "acconts" (
  "id" uuid PRIMARY KEY NOT NULL,
  "holderId" uuid,
  "balance" decimal NOT NULL,
  "currency" USD,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "Transaction" (
  "transactionId" uuid,
  "from" uuid,
  "to" uuid,
  "amount" decimal,
  "created_at" timestamptz DEFAULT (now())
);

ALTER TABLE "acconts" ADD FOREIGN KEY ("holderId") REFERENCES "users" ("id");

ALTER TABLE "Transaction" ADD FOREIGN KEY ("from") REFERENCES "acconts" ("id");

ALTER TABLE "Transaction" ADD FOREIGN KEY ("to") REFERENCES "acconts" ("id");

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "acconts" ("holderId");

CREATE INDEX ON "Transaction" ("from");

CREATE INDEX ON "Transaction" ("to");

CREATE INDEX ON "Transaction" ("transactionId");
