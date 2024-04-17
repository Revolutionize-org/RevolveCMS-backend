package model

type Role struct {
	tableName struct{} `pg:"role"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
}
