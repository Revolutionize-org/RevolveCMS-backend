CREATE TABLE website (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    theme_id UUID NOT NULL,
    header_id UUID NOT NULL,
    page_id UUID NOT NULL,
    footer_id UUID NOT NULL,
    CONSTRAINT fk_header_id FOREIGN KEY (header_id) REFERENCES header (id) ON DELETE CASCADE,
    CONSTRAINT fk_page_id FOREIGN KEY (page_id) REFERENCES page (id) ON DELETE CASCADE,
    CONSTRAINT fk_footer_id FOREIGN KEY (footer_id) REFERENCES footer (id) ON DELETE CASCADE,
    CONSTRAINT fk_theme_id FOREIGN KEY (theme_id) REFERENCES theme (id)
);
