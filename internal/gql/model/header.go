package model

type Header struct {
	ID        string `json:"id" validate:"uuid4"`
	Name      string `json:"name"`
	Data      string `json:"data"`
	CreatedAt string `json:"created_at" pg:"-"`
	UpdatedAt string `json:"updated_at"`
	WebsiteID string `json:"website_id"`
}

type HeaderInput struct {
	ID   *string `json:"id,omitempty" validate:"uuid4"`
	Name string  `json:"name"`
	Data string  `json:"data"`
}
