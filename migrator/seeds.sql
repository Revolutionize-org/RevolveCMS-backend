INSERT INTO role(
    id,
    name
) VALUES
    ('d7de28aa-5028-4bee-8361-7d630d86da54', 'admin'),
    ('d44e3f29-0ab5-40d4-b5d0-1e41c3cc59d3', 'user');

INSERT INTO theme(
    id,
    name
) VALUES 
    ('01faca9d-fa3c-4c2f-bb98-5fd4de0f9cc3', 'white'),
    ('74d0abe1-8033-4441-9e91-ef238bf1eadd', 'dark');

INSERT INTO website(
    id,
    name,
    theme_id
) VALUES 
    ('45955517-30ee-4310-b253-d0cd677cc92e', 'Revolutionize', '01faca9d-fa3c-4c2f-bb98-5fd4de0f9cc3');

INSERT INTO header (
    id,
    name,
    data,
    updated_at,
    website_id
) VALUES (
    '61759ae3-f8eb-4b51-a96c-b5fa143a6d85',
    'Revolutionize Header',
    '{"text": "My Text Header"}',
    CURRENT_TIMESTAMP,
    '45955517-30ee-4310-b253-d0cd677cc92e'
);

INSERT INTO footer (
    id,
    name,
    data,
    updated_at,
    website_id
) VALUES (
    'e397511b-1cba-414f-a3ea-4173ceef171d',
    'Revolutionize Footer',
    '{"text": "My Text Footer"}',
    CURRENT_TIMESTAMP,
    '45955517-30ee-4310-b253-d0cd677cc92e'
);

INSERT INTO page (
    id,
    name,
    slug,
    data,
    updated_at,
    website_id
) VALUES (
    'e397511b-1cba-414f-a3ea-4173ceef171d',
    'Revolutionize Page',
    'dashboard',
    '{"text": "My Text Page"}',
    CURRENT_TIMESTAMP,
    '45955517-30ee-4310-b253-d0cd677cc92e'
);