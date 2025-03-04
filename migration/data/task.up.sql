CREATE TABLE IF NOT EXISTS tasks (
                                     id BIGSERIAL PRIMARY KEY,
                                     name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    order_index INTEGER NOT NULL
    );

COPY tasks (id, name, description, created_at, order_index)
    FROM '/docker-entrypoint-initdb.d/data'
    DELIMITER ',' CSV HEADER;

CREATE INDEX idx_tasks_order_index ON tasks (order_index);

CREATE INDEX idx_tasks_created_at ON tasks (created_at);