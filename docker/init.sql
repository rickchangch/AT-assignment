CREATE TABLE IF NOT EXISTS users (
    acct VARCHAR(32) PRIMARY KEY,
    pwd VARCHAR(72) NOT NULL,
    fullname VARCHAR(64) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
)
