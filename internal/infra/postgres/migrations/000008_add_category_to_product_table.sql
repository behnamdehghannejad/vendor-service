-- +migrate Up
ALTER TABLE products
    ADD COLUMN category_id BIGINT;

ALTER TABLE products
    ADD CONSTRAINT fk_products_categories
        FOREIGN KEY (category_id) REFERENCES categories(id);

-- +migrate Down

ALTER TABLE products
DROP CONSTRAINT IF EXISTS fk_products_categories;

ALTER TABLE products
DROP COLUMN IF EXISTS category_id;