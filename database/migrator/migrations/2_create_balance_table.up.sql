DROP TABLE IF EXISTS balances;
CREATE TABLE balances (
    id VARCHAR(60) PRIMARY KEY,
    total decimal(10,2) NOT NULL,
    debit_avarage decimal(10,2) NOT NULL,
    credit_avarage decimal(10,2) NOT NULL,
    transactions VARCHAR(60),
    user_id VARCHAR(60) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);