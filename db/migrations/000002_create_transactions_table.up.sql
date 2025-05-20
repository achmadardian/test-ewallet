CREATE TYPE transaction_status AS ENUM (
    'PENDING',
    'SUCCESS',
    'FAILED'
);

CREATE TYPE transaction_type AS ENUM (
    'DEBIT',
    'CREDIT'
);

CREATE TABLE IF NOT EXISTS transactions (
    id UUID PRIMARY KEY NOT NULL,
    status transaction_status NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    transaction_type transaction_type NOT NULL,
    amount BIGINT NOT NULL,
    remarks VARCHAR(50) NULL,
    balance_before BIGINT NOT NULL,
    balance_after BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP
);