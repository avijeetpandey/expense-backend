CREATE TABLE "expenses" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "tag" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "expenses" ("name");
