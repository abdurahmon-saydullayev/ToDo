CREATE TABLE IF NOT EXISTS todo(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP default (NOW()),
    updated_at TIMESTAMP default (NOW()),
    completed boolean default  false
);