package model

type Website struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	ThemeID string `json:"theme_id"`
	// Header *Header `json:"header,omitempty"`
	// Pages  []*Page `json:"pages"`
	// Footer *Footer `json:"footer,omitempty"`
}
