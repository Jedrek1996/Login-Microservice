CREATE TABLE "user_details" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "user_name" varchar NOT NULL UNIQUE,
  "user_password" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "mobile" int NOT NULL UNIQUE,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_cookies" (
  "id" SERIAL PRIMARY KEY,
  "user_name" varchar NOT NULL UNIQUE,
  "cookie_id" int NOT NULL UNIQUE,
  "expires_at" TIMESTAMP NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
  FOREIGN KEY ("user_name") REFERENCES "user_details" ("user_name")
);

CREATE TABLE "customer_address" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "user_id" int,
  "address_id" int
);

CREATE TABLE "address" (
  "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
  "unit_number" varchar NOT NULL,
  "address_line1" varchar NOT NULL,
  "address_line2" varchar NOT NULL,
  "postal_code" int NOT NULL
);

CREATE INDEX ON "user_details" ("id");
CREATE INDEX ON "user_cookies" ("user_name");
CREATE INDEX ON "user_cookies" ("cookie_id");
CREATE INDEX ON "address" ("id");
CREATE INDEX ON "address" ("postal_code");

ALTER TABLE "customer_address" ADD FOREIGN KEY ("user_id") REFERENCES "user_details" ("id");
ALTER TABLE "customer_address" ADD FOREIGN KEY ("address_id") REFERENCES "address" ("id");