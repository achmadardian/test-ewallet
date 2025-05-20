CREATE TABLE IF NOT EXISTS topups (
    id UUID PRIMARY KEY NOT NULL,
    transaction_id UUID NOT NULL REFERENCES transactions(id),
    amount BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);