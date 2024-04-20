CREATE TABLE footer(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    data TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    website_id UUID NOT NULL UNIQUE,
    CONSTRAINT fk_website_id FOREIGN KEY (website_id) REFERENCES website(id)
);