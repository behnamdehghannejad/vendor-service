-- +migrate Up

CREATE TABLE vendors (
    id BIGSERIAL PRIMARY KEY,

    code VARCHAR(50) NOT NULL,
    name VARCHAR(200) NOT NULL,

    email VARCHAR(100),
    phone VARCHAR(20),
    address VARCHAR(500),

    active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT uq_vendor_code UNIQUE (code)
);

-- +migrate Down

DROP TABLE IF EXISTS vendors;