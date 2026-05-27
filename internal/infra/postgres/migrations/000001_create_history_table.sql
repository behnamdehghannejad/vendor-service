-- +migrate Up
CREATE TABLE history (
    id BIGSERIAL PRIMARY KEY,

    order_id UUID NOT NULL,
    payment_id UUID NOT NULL,

    quantity INT NOT NULL,
    product_id INT NOT NULL,
    vendor_id INT NOT NULL,

    status VARCHAR(50) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


-- +migrate Down
DROP TABLE IF EXISTS history;