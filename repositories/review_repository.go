package repositories

import (
	"go_api_echo/db"
	"go_api_echo/queries"
	"go_api_echo/transport"
)

func CreateReview(data transport.CreateReviewResponse) error {
	createReviewQuery := queries.CreateReviewByIdQuery(data)

	_, err := db.Mysql(createReviewQuery)
	if err != nil {
		return err
	}
	return nil
}