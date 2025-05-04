CREATE TABLE posts(
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    parent_post_id INT,
    user_id INT NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ
);