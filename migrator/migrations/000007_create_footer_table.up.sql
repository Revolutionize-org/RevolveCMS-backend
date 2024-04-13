CREATE TABLE footer(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    data TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ,

    website_id UUID NOT NULL,
    CONSTRAINT fk_website_id FOREIGN KEY (website_id) REFERENCES website(id)
);