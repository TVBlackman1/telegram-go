CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(64),
    chat_id BIGINT NOT NULL,
    state_id UUID
);