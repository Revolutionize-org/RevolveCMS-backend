package model

type User struct {
	tableName    struct{} `pg:"users"`
	ID           string   `json:"id" validate:"uuid4"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	PasswordHash string   `json:"password_hash"`
	CreatedAt    string   `json:"created_at" pg:"-"`
	RoleID       string   `json:"role_id" validate:"uuid4"`
	WebsiteID    string   `json:"website_id" validate:"uuid4"`
}

type UserInfo struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=5,max=32"`
}
