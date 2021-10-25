package entity

type Movie struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"-"`
	Title string `gorm:"" json:"Title"`
	Year string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type string `json:"Type"`
	Poster string `json:"Poster"`
}