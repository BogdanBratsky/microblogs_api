CREATE TABLE follows(
    follower_id INT NOT NULL REFERENCES users(id),
    followee_id INT NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ DEFAULT now(),
    PRIMARY KEY (follower_id, followee_id)
);