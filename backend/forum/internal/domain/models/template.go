package models

type Template struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Photos []Photo `json:"photos"`
	HTML   string  `json:"html"`
}
