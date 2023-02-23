package model

type MovieInput struct {
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
	Length int    `json:"length"`
}

type Movie struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Year   int    `json:"year"`
	Length int    `json:"length"`
}

type MovieID struct {
	ID int `json:"id"`
}

type ForceGetMovies struct {
	Force bool `json:"force"`
}

type MovieList struct {
	Movies []Movie `json:"movies"`
}
