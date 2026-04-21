-- create tables

-- Accounts
CREATE TABLE accounts
(
    id  BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    balance BIGINT NOT NULL DEFAULT 0, -- stored in cents
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Transactions
CREATE TABLE Transactions(
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    amount  BIGINT NOT NULL,
    description TEXT,
    category INT NOT NULL,
    transaction_type INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_account
        FOREIGN KEY (account_id)
        REFERENCES accounts(id)
        ON DELETE CASCADE
);

-- Add Indexes
CREATE INDEX idx_transactions_account_id
ON transactions(account_id);

CREATE INDEX idx_transactions_created_at
ON transactions(created_at);