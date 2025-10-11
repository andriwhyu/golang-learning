CREATE TABLE IF NOT EXISTS comments(
    id SERIAL PRIMARY KEY,
    username VARCHAR(50),
    content TEXT NOT NULL
);