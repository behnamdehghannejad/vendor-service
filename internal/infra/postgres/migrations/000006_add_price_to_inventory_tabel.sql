-- +migrate Up

ALTER TABLE inventories
ADD COLUMN price INTEGER NOT NULL DEFAULT 0,
ADD CONSTRAINT inventories_price_check
CHECK (price >= 0);

-- +migrate Down

ALTER TABLE inventories
DROP CONSTRAINT IF EXISTS inventories_price_check,
DROP COLUMN IF EXISTS price;