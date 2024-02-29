CREATE TABLE goods (
    id SERIAL PRIMARY KEY,
    project_id INT,
    name VARCHAR(255),
    description VARCHAR(255),
    priority INT,
    removed bool,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);