CREATE TABLE IF NOT EXISTS users (
  id   int PRIMARY KEY,
  name text      NOT NULL,
  email  text NOT NULL
);

CREATE TABLE IF NOT EXISTS urls (
  id   int PRIMARY KEY,
  url text      NOT NULL,
  shorturl  text NOT NULL
);