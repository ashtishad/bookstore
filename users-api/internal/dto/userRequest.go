package dto

type UserRequest struct {
	Name        string `json:"name" db:"name" binding:"required"`
	Gender      string `json:"gender" db:"gender" binding:"required"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth" binding:"required"`
	Email       string `json:"email" db:"email" binding:"required,email"`
	City        string `json:"city" db:"city" binding:"required"`
}

type UserUpdateRequest struct {
	Id     int64  `json:"id" db:"id" binding:"required"`
	Name   string `json:"name" db:"name" binding:"required"`
	Email  string `json:"email" db:"email" binding:"required,email"`
	City   string `json:"city" db:"city" binding:"required"`
	Status *int8  `json:"status" db:"status" binding:"required,oneof=0 1"`
}

func (uu *UserUpdateRequest) SetId(id int64) {
	uu.Id = id
}
