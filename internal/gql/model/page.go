package model

type Page struct {
	ID        string `json:"id" validate:"uuid4"`
	Name      string `json:"name" validate:"omitempty,min=2,max=32"`
	Slug      string `json:"slug" validate:"omitempty,min=2,max=32"`
	Data      string `json:"data"`
	CreatedAt string `json:"created_at" pg:"-"`
	UpdatedAt string `json:"updated_at" validate:"omitempty,timezone"`
	WebsiteID string `json:"website_id" validate:"omitempty,uuid4"`
}

type PageInput struct {
	ID   *string `json:"id,omitempty" validate:"omitempty,uuid4"`
	Name string  `json:"name" validate:"omitempty,min=2,max=32"`
	Slug string  `json:"slug" validate:"omitempty,min=2,max=32,excludesall= "`
	Data string  `json:"data"`
}
