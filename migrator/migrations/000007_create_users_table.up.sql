CREATE TABLE users(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    role_id UUID NOT NULL,
    
    CONSTRAINT fk_role_id FOREIGN KEY (role_id) REFERENCES role(id)
);