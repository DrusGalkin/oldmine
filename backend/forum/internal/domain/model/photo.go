package model

type Photo struct {
	ID     int    `json:"id"`
	Url    string `json:"url"`
	Index  int    `json:"index"`
	UserID int    `json:"user_id"`
}
