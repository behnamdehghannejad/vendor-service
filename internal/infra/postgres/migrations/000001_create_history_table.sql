-- +migrate Up
CREATE TABLE histories (
    id VARCHAR(36) PRIMARY KEY,

    reserved INT NOT NULL CHECK (quantity > 0),

    product_id INT NOT NULL,
    vendor_id INT NOT NULL,

    status VARCHAR(50) NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE IF EXISTS histories;