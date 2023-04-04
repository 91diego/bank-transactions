DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions (
    id VARCHAR(60) PRIMARY KEY,
    transaction_date TIMESTAMP NOT NULL,
    transaction_amount decimal(10,2) NOT NULL,
    transaction_type VARCHAR(30) NOT NULL,
    balance_id VARCHAR(60) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);