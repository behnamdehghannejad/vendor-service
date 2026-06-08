-- +migrate Up
CREATE TABLE categories (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(200) NOT NULL UNIQUE,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  path VARCHAR(1000),
  parent_id BIGINT,

  CONSTRAINT fk_categories_parent FOREIGN KEY (parent_id) REFERENCES categories(id)
);

-- +migrate Down
DROP TABLE IF EXISTS categories;