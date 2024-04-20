package model

type Website struct {
	ID      string  `json:"id" validate:"uuid4"`
	Name    string  `json:"name"`
	ThemeID string  `json:"theme_id" validate:"uuid4"`
	Header  *Header `json:"header,omitempty" pg:"-"`
	Pages   []*Page `json:"pages,omitempty" pg:"-"`
	Footer  *Footer `json:"footer,omitempty" pg:"-"`
}
