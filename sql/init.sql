CREATE TABLE users (
    acct VARCHAR(32) PRIMARY KEY,
    pwd VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP
)
