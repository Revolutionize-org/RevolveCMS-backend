package model

type Footer struct {
	ID        string `json:"id" validate:"uuid4"`
	Name      string `json:"name"`
	Data      string `json:"data"`
	CreatedAt string `json:"created_at" pg:"-"`
	UpdatedAt string `json:"updated_at" validate:"timezone"`
	WebsiteID string `json:"website_id" validate:"uuid4"`
}

type FooterInput struct {
	ID   *string `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name string  `json:"name" validate:"omitempty,min=2,max=32"`
	Data string  `json:"data"`
}
