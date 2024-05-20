CREATE TYPE gender AS ENUM ('male', 'female', 'other');


CREATE TABLE IF NOT EXISTS customers (
  id uuid PRIMARY KEY,
  external_id VARCHAR,
  first_name varchar,
  last_name varchar,
  age INTEGER,
  phone VARCHAR[],
  mail VARCHAR,
  birthday VARCHAR,
  sex gender,
  created_at TIMESTAMP NOT NULL default 'now()',
  updated_at TIMESTAMP);

