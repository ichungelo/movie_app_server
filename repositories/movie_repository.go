package repositories

import (
	"go_api_echo/db"
	"go_api_echo/entities"
	"go_api_echo/queries"
)

func ReadAllMovies() (data []entities.MovieEntity, err error) {
	readAllMoviesQuery := queries.ReadAllMoviesQuery()
	queryResult := db.Mysql(readAllMoviesQuery)
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
	return
}
