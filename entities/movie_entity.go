package entities

type MovieEntity struct {
	MovieId     int    `json:"movie_id"`
	Title       string `json:"title"`
	ReleaseYear int    `json:"release_year"`
	Production  string `json:"production"`
	Overview    string `json:"overview"`
}
