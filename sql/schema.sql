CREATE TABLE feedback (
    id SERIAL PRIMARY KEY,
    rating INTEGER NOT NULL CHECK (rating >= 0 AND rating <= 5),
    description VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    id_figure INTEGER,
    id_user INTEGER
);

CREATE TABLE favorite (
    id_user   INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    id_figure INT NOT NULL REFERENCES figure(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id_user, id_figure)
);