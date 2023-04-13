CREATE TABLE IF NOT EXISTS users (
  id   INTEGER NOT NULL  PRIMARY KEY AUTOINCREMENT,
  name text      NOT NULL,
  email  text NOT NULL
);

CREATE TABLE IF NOT EXISTS urls (
  id  INTEGER PRIMARY KEY AUTOINCREMENT,
  url text      NOT NULL,
  shorturl  text NOT NULL,
  userid int,
  createdate date,
  FOREIGN KEY(userid) REFERENCES users(id)
);