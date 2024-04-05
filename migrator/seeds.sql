INSERT INTO role(
    id,
    name
) VALUES
    ('d7de28aa-5028-4bee-8361-7d630d86da54', 'admin'),
    ('d44e3f29-0ab5-40d4-b5d0-1e41c3cc59d3', 'user');

INSERT INTO users(
    id,
    name,
    email,
    password_hash,
    role_id
) VALUES
    ('c7e4b9d6-2b3b-4a0e-8d6d-6f5b0f9e0f9e', 'admin', 'aKu1X@example.com', 'admin', 'd7de28aa-5028-4bee-8361-7d630d86da54');