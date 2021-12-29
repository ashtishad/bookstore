package dto

type UserRequest struct {
	Name        string `json:"name" db:"name" binding:"required"`
	Gender      string `json:"gender" db:"gender" binding:"required"`
	DateOfBirth string `json:"date_of_birth" db:"date_of_birth" binding:"required"`
	Email       string `json:"email" db:"email" binding:"required,email"`
	City        string `json:"city" db:"city" binding:"required"`
}

// TODO: Add validation for UserRequest

func (u *UserRequest) Validate() error {
	return nil
}
