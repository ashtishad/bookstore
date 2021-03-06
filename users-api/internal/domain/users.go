package domain

import (
	"github.com/ashtishad/bookstore/lib"
	"github.com/ashtishad/bookstore/users-api/internal/dto"
	"time"
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
	Save(User) (*User, lib.RestErr)
	Update(User) (*User, lib.RestErr)
	Delete(id int64) lib.RestErr
	FindByName(name string) (*[]User, lib.RestErr)

	//FindAll() ([]*User, lib.RestErr)
}

// CREATE : User -> DTO , vice versa transformation

func NewUser(req dto.UserRequest) User {
	return User{
		Name:        req.Name,
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
		Email:       req.Email,
		City:        req.City,
		DateCreated: time.Now().Format(lib.DbDateLayout),
		Status:      1,
	}
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

// NewUpdateUser has only fields that are updatable
func NewUpdateUser(req dto.UserUpdateRequest) User {
	return User{
		Id:     req.Id,
		Name:   req.Name,
		Email:  req.Email,
		City:   req.City,
		Status: *req.Status,
	}
}

// ToUserSearchRespDTO is used to transform the User to UserSearchResponse
func (u User) ToUserSearchRespDTO() dto.UserSearchResponse {
	return dto.UserSearchResponse{
		Id:          u.Id,
		Name:        u.Name,
		Email:       u.Email,
		DateCreated: u.DateCreated,
		Status:      u.Status,
	}
}
