package domain

import (
	"github.com/ashtishad/bookstore/lib"
	"github.com/ashtishad/bookstore/users-api/internal/dto"
)

type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
	Email       string `json:"email"`
	City        string `json:"city"`
	DateCreated string `json:"date_created"`
	Status      int8   `json:"status"`
}

// UserRepository is the interface that wraps the basic CRUD functions
type UserRepository interface {
	FindById(id int64) (*User, lib.RestErr)

	//Create(user *UserResponse) (*UserResponse, lib.RestErr)
	//Search(name string) (*UserResponse, lib.RestErr)
	//FindAll() ([]*UserResponse, lib.RestErr)
	//Update(user *UserResponse) (*UserResponse, lib.RestErr)
	//Delete(id int64) lib.RestErr
}

func (u User) ToUserRespDTO() dto.UserResponse {
	return dto.UserResponse{
		Id:          u.Id,
		Name:        u.Name,
		Gender:      u.Gender,
		DateOfBirth: u.DateOfBirth,
		Email:       u.Email,
		City:        u.City,
		DateCreated: u.DateCreated,
		Status:      u.Status,
	}
}
