CREATE INDEX idx_header_website_id ON header(website_id);
CREATE INDEX idx_footer_website_id ON footer(website_id);
CREATE INDEX idx_page_website_id ON page(website_id);
CREATE INDEX idx_users_website_id ON users(website_id);
CREATE INDEX idx_role_id ON users(role_id);
CREATE INDEX idx_theme_id ON website(theme_id);
