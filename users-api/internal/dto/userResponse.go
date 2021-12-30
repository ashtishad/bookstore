package dto

type UserResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
	Email       string `json:"email"`
	City        string `json:"city"`
	DateCreated string `json:"date_created"`
	Status      int8   `json:"status"`
}

type UserSearchResponse struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      int8   `json:"status"`
}
