package usecases

import (
	"go_api_echo/entities"
	"go_api_echo/transport"
)

func JwtPayloadUsecase(data *entities.UserLoginEntity) entities.JwtPayloadEntity {
	payload := entities.JwtPayloadEntity{
		UserId: data.UserId,
		Username: data.Username,
		Email: data.Email,
		FirstName: data.FirstName,
		LastName: data.LastName,
	}

	return payload
}

func PostReviewUsecase(userId string, movieId int, review string) transport.CreatePostReviewRequest {
	data := transport.CreatePostReviewRequest{
		UserId: userId,
		MovieId: movieId,
		Review: review,
	}

	return data
}

func PutReviewUsecase(reviewId int, userId string, movieId int, review string) transport.PutReviewRequest {
	data := transport.PutReviewRequest{
		ReviewId: reviewId,
		UserId: userId,
		MovieId: movieId,
		Review: review,
	}

	return data
}

func DeleteReviewUsecase(reviewId int ,userId string, movieId int) transport.DeleteReviewRequest {
	data := transport.DeleteReviewRequest{
		ReviewId: reviewId,
		UserId: userId,
		MovieId: movieId,
	}

	return data
}