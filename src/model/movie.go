package model

type Movie struct {
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
	Length int    `json:"length"`
}

type MovieID struct {
	ID int `json:"id"`
}

type MovieList struct {
	Movies []Movie `json:"movies"`
}
