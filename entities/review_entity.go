package entities

type ReviewEntity struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	MovieId  int    `json:"movie_id"`
	Review   string `json:"review"`
}
