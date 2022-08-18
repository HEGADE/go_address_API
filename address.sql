BEGIN;
CREATE TABLE "mst_state"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    CONSTRAINT "mst_state_name_ukey" UNIQUE ("name")
);

CREATE TABLE "mst_city"(
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "mst_state_id" INTEGER NOT  NULL,
    CONSTRAINT "mst_city_mst_state_state_fkey" FOREIGN KEY ("mst_state_id") REFERENCES mst_state("id") ON DELETE CASCADE
);

COMMIT;
