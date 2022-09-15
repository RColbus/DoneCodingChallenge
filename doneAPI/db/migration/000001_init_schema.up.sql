CREATE TABLE "registrations" (
                                 "id" bigserial PRIMARY KEY,
                                 "first_name" varchar NOT NULL,
                                 "middle_name" varchar NOT NULL,
                                 "last_name" varchar NOT NULL,
                                 "phone_number" varchar NOT NULL,
                                 "email_address" varchar NOT NULL,
                                 "street_address_line1" varchar NOT NULL,
                                 "street_address_line2" varchar NOT NULL,
                                 "city" varchar NOT NULL,
                                 "state" varchar NOT NULL,
                                 "zip_code" varchar NOT NULL,
                                 "photo_path" varchar NOT NULL,
                                 "birth_date" varchar NOT NULL,
                                 "appointment_date" varchar NOT NULL,
                                 "appointment_time" varchar NOT NULL,
                                 "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "registrations" ("last_name");

CREATE INDEX ON "registrations" ("birth_date");

CREATE INDEX ON "registrations" ("first_name", "last_name");

CREATE INDEX ON "registrations" ("zip_code");
