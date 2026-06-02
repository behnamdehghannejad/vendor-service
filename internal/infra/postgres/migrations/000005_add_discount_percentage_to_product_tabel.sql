-- +migrate Up

ALTER TABLE inventories
ADD COLUMN discount_percentage NUMERIC(5,2) NOT NULL DEFAULT 0.00,
ADD CONSTRAINT products_discount_percentage_check
CHECK (discount_percentage >= 0 AND discount_percentage <= 100);

-- +migrate Down

ALTER TABLE inventories
DROP CONSTRAINT IF EXISTS inventories_discount_percentage_check,
DROP COLUMN IF EXISTS discount_percentage;