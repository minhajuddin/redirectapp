CREATE TABLE redirects
(
  id SERIAL,
  host VARCHAR(255) NOT NULL UNIQUE,
  rules TEXT
)
