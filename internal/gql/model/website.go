package model

type Website struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	ThemeID string  `json:"theme_id"`
	Header  *Header `json:"header,omitempty" pg:"-"`
	Pages   []*Page `json:"pages" pg:"-"`
	Footer  *Footer `json:"footer,omitempty" pg:"-"`
}
