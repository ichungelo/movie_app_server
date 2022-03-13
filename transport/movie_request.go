package transport

type GetMovieRequest struct {
	MovieId      string `json:"movie_id" validate:"required"`
}