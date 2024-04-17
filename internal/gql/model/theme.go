package model

type Theme struct {
	tableName struct{} `pg:"theme"`
	ID        string   `json:"id"`
	Name      string   `json:"name"`
}
