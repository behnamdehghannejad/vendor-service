-- +migrate Up

CREATE TABLE inventories (
    id BIGSERIAL PRIMARY KEY,

    vendor_id INT NOT NULL,
    product_id INT NOT NULL,

    quantity INT NOT NULL CHECK (quantity >= 0),
    reserved INT NOT NULL DEFAULT 0 CHECK (reserved >= 0),

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT uq_inventory_vendor_product UNIQUE (vendor_id, product_id),
    CONSTRAINT chk_inventory_reserved_leq_quantity CHECK (reserved <= quantity),

    CONSTRAINT fk_inventories_vendor
        FOREIGN KEY (vendor_id) REFERENCES vendors(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_inventories_product
        FOREIGN KEY (product_id) REFERENCES products(id)
        ON DELETE CASCADE
);

CREATE INDEX idx_inventories_vendor_id ON inventories(vendor_id);
CREATE INDEX idx_inventories_product_id ON inventories(product_id);

-- +migrate Down

DROP TABLE IF EXISTS inventories;