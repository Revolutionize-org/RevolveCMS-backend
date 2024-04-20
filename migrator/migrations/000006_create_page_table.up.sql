CREATE TABLE page(
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    slug TEXT NOT NULL UNIQUE,
    data TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    website_id UUID NOT NULL,
    CONSTRAINT fk_website_id FOREIGN KEY (website_id) REFERENCES website(id)
);