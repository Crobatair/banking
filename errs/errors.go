package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func NewNotFoundError(message string) *AppError {
	return &AppError{http.StatusNotFound, message}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{http.StatusInternalServerError, message}
}
