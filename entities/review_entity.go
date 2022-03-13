package entities

type ReviewEntity struct {
	ReviewId int `json:"review_id"`
	Username string `json:"username"`
	Review   string `json:"review"`
}
