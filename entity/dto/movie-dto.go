package dto

type MovieResponse struct {
	Search []MovieDTO `json:"Search"`
}

type MovieDTO struct {
	Title string `json:"Title"`
	Year string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type string `json:"Type"`
	Poster string `json:"Poster"`
}