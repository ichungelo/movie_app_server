package repositories

import (
	"go_api_echo/db"
	"go_api_echo/entities"
	"go_api_echo/queries"
)

func ReadAllMovies() ([]entities.MovieEntity, error) {
	data := []entities.MovieEntity{}
	readAllMoviesQuery := queries.ReadAllMoviesQuery()
	queryResult, err := db.Mysql(readAllMoviesQuery)
	if err != nil {
		return nil, err
	}

	for queryResult.Next() {
		movie := new(entities.MovieEntity)
		err := queryResult.Scan(&movie.MovieId, &movie.Title, &movie.ReleaseYear, &movie.Production, &movie.Overview)
		if err != nil {
			return nil, err
		}
		data = append(data, *movie)
	}

	defer queryResult.Close()
	return data, nil
}

func ReadMovieById(param int) (*entities.MovieReviewEntity, error)  {
	movie := entities.MovieReviewEntity{}
	reviews := []entities.ReviewEntity{}

	movieQuery, reviewQuery := queries.ReadMovieByIdQuery(param)
	movieQueryResult, err := db.Mysql(movieQuery)
	if err != nil {
		return nil, err
	}

	for movieQueryResult.Next() {
		err := movieQueryResult.Scan(&movie.MovieId, &movie.Title, &movie.ReleaseYear, &movie.Production, &movie.Overview)
		if err != nil {
			return nil, err
		}

	}

	reviewQueryResult, err := db.Mysql(reviewQuery)
	if err != nil {
		return nil, err
	}

	for reviewQueryResult.Next() {
		review := new(entities.ReviewEntity)
		err := reviewQueryResult.Scan( &review.ReviewId, &review.Username, &review.Review)
		if err != nil {
			return nil, err
		}
		reviews = append(reviews, *review)
	}

	data := entities.MovieReviewEntity{
		MovieId: movie.MovieId,
		Title: movie.Title,
		ReleaseYear: movie.ReleaseYear,
		Production: movie.Production,
		Overview: movie.Overview,
		Reviews: reviews,
	}

	return &data, nil
}
