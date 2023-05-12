CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name text NOT NULL,
  password text NOT NULL,
  apikey text NOT NULL
);

CREATE TABLE IF NOT EXISTS urls (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  url text NOT NULL,
  shorturl text NOT NULL,
  userid INTEGER,
  createdate INTEGER,
  FOREIGN KEY(userid) REFERENCES users(id)
);