CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    cpf VARCHAR(11) NOT NULL UNIQUE,
    name VARCHAR(150) NOT NULL,
    email VARCHAR(320) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role INT DEFAULT 1,
    creation_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deactivated_at TIMESTAMP
);

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