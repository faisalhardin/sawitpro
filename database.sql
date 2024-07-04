-- This is the SQL script that will be used to initialize the database schema.
-- We will evaluate you based on how well you design your database.
-- 1. How you design the tables.
-- 2. How you choose the data types and keys.
-- 3. How you name the fields.
-- In this assignment we will use PostgreSQL as the database.

-- This is test table. Remove this table and replace with your own tables. 
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE swp_mst_estate (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL DEFAULT uuid_generate_v4(),
    width INTEGER NOT NULL,
    length INTEGER NOT NULL,
	create_time TIMESTAMPTZ NOT NULL,
    update_time TIMESTAMPTZ NOT NULL,
    delete_time TIMESTAMPTZ,
    UNIQUE (uuid)
);

CREATE TABLE public.swp_trx_tree_estate (
    id SERIAL PRIMARY KEY,
	uuid UUID DEFAULT uuid_generate_v4() UNIQUE,
    position_x INT NOT NULL,
    position_y INT NOT NULL,
    height INT NOT NULL,
    id_mst_estate BIGINT NOT NULL,
    create_time TIMESTAMPTZ NOT NULL,
    update_time TIMESTAMPTZ NOT NULL,
    delete_time TIMESTAMPTZ
);