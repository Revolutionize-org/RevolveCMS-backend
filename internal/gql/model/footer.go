package model

type Footer struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Data      string `json:"data"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	WebsiteID string `json:"website_id"`
}

type FooterInput struct {
	ID   *string `json:"id,omitempty"`
	Name string  `json:"name"`
	Data string  `json:"data"`
}
