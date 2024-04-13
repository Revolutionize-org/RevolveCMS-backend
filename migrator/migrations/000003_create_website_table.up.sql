CREATE TABLE website (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    
    theme_id UUID NOT NULL,
    CONSTRAINT fk_theme_id FOREIGN KEY (theme_id) REFERENCES theme (id)
);
