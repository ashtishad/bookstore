package lib

import (
	"net/http"
)

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"statusCode,omitempty"`
}

func (e *AppError) AsMessage() string {
	return e.Message
}

func NewNotFoundError(m string) *AppError {
	return &AppError{
		Message: m,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(m string) *AppError {
	return &AppError{
		Message: m,
		Code:    http.StatusInternalServerError,
	}
}

func NewInternalServerError(m string) *AppError {
	return &AppError{
		Message: m,
		Code:    http.StatusInternalServerError,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}
