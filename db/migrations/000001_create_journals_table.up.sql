CREATE TABLE IF NOT EXISTS journals (
  id serial PRIMARY KEY,
  name VARCHAR,
  created_at TIMESTAMP WITHOUT TIME ZONE,
  updated_at TIMESTAMP WITHOUT TIME ZONE
);
