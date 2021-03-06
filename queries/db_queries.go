package queries

import (
	"fmt"
	"go_api_echo/transport"
	"go_api_echo/utils"
	"time"
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
	queryMovie = fmt.Sprintf("SELECT * FROM movies WHERE movie_id=%d", param)
	queryReview = fmt.Sprintf("SELECT reviews.review_id, users.username, reviews.review FROM reviews INNER JOIN users ON users.user_id=reviews.user_id WHERE movie_id=%d AND is_delete=false", param)

	return
}

func CreateReviewByIdQuery(data transport.CreatePostReviewRequest) (query string) {
	query = fmt.Sprintf("INSERT INTO reviews (user_id, movie_id, review) VALUES ('%s', %d, '%s')", data.UserId, data.MovieId, data.Review )
	return
}

func PutReviewQuery(data transport.PutReviewRequest) (query string) {
	query = fmt.Sprintf("UPDATE reviews SET review='%s' WHERE review_id=%d AND user_id='%s' AND movie_id=%d", data.Review, data.ReviewId, data.UserId, data.MovieId)
	return
}

func DeleteReviewQuery(data transport.DeleteReviewRequest) (query string) {
	time := time.Now().Format("2006-01-02 15:04:05")
	query = fmt.Sprintf("UPDATE reviews SET is_delete=true, deleted_at='%s' WHERE review_id=%d AND user_id='%s' AND movie_id=%d",time, data.ReviewId, data.UserId, data.MovieId)
	return
}
