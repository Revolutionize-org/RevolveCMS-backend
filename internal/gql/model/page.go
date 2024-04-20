package model

type Page struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Data      string `json:"data"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	WebsiteID string `json:"website_id"`
}

type PageInput struct {
	ID   *string `json:"id,omitempty"`
	Name string  `json:"name"`
	Slug string  `json:"slug" validate:"excludesall= "`
	Data string  `json:"data"`
}
