CREATE TABLE IF NOT EXISTS transfers (
    id UUID PRIMARY KEY NOT NULL,
    sender_transaction_id UUID NOT NULL REFERENCES transactions(id),
    receiver_transaction_id UUID NOT NULL REFERENCES transactions(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
)