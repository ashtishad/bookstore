CREATE TABLE "users" (
                    "id" bigserial PRIMARY KEY NOT NULL,
                    "name" varchar(50) NOT NULL,
                    "gender" varchar(10) NOT NULL,
                    "date_of_birth" date NOT NULL,
                    "email" varchar(40) NOT NULL,
                    "city" varchar(30) NOT NULL,
                    "date_created" timestamp NOT NULL DEFAULT 'now()',
                    "status" SMALLINT NOT NULL DEFAULT 1
);