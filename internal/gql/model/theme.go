package model

type Theme struct {
	ID   string `json:"id" validate:"uuid4"`
	Name string `json:"name"`
}
