package errs

import "net/http"

type AppError struct {
	Code    int    `json:"code" ,omitempty:"true"`
	Message string `json:"message"`
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
func NewNotFoundError(message string) *AppError {
	return &AppError{http.StatusNotFound, message}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{http.StatusInternalServerError, message}
}

func NewBadRequestError(message string) *AppError {
	return &AppError{http.StatusBadRequest, message}
}

func NewValidationError(message string) *AppError {
	return &AppError{http.StatusUnprocessableEntity, message}
}
