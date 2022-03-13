package queries

import (
	"fmt"
	"go_api_echo/transport"
	"go_api_echo/utils"
)

func CreateUserQuery(data transport.RegisterRequest) (query string) {
	userId := utils.IdGenerator(36)
	hashedPassword, _ := utils.HashPassword(data.Password)

	query = fmt.Sprintf("INSERT INTO users (user_id, username, email, first_name, last_name, password) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')", userId, data.Username, data.Email, data.FirstName, data.LastName, hashedPassword)
	return
}

func ReadUserQuery(data transport.LoginRequest) (query string) {
	query = fmt.Sprintf("SELECT user_id, username, email, first_name, last_name, password FROM users WHERE username='%s'", data.Username)
	return
}

func ReadAllMoviesQuery() (query string) {
	query = "SELECT * FROM movies"
	return
}

func ReadMovieByIdQuery(param int) (queryMovie string, queryReview string) {
	queryMovie = fmt.Sprintf("SELECT * FROM movies WHERE movie_id='%d'", param)
	queryReview = fmt.Sprintf("SELECT * FROM reviews WHERE movie_id=%d", param)
	return
}

// func CreateReviewByIdQuery(userId string, movieId string) (query string) {
// 	query = fmt.Sprintf("INSERT INTO reviews (user_id, movie_id, review) VALUES ('%s', %s, '%s')", userId, mov )
// }

// func MoviesDbQueries(id string) {
// 	GetAllMovies := "SELECT * FROM movies WHERE is_deleted = false"
// 	GetMovieById := "SELECT movies.movie_id, "
// }
