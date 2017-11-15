-- +mig Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE name_ref (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name text,
    middle_name text,
    last_name text,
    suffix text
);

CREATE TABLE human (
    votem_uhid uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name_ref_id uuid NOT NULL REFERENCES name_ref(id) ON DELETE restrict ON UPDATE cascade,
    primary_email text NOT NULL UNIQUE,
    primary_phone text,
    username text UNIQUE,
    password_hash text NOT NULL
);

-- +mig Down
DROP TABLE human;
DROP TABLE name_ref;
