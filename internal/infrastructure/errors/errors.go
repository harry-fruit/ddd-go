package apperror

import (
	"errors"
)

type InternalServerError struct {
	Err error
}

func NewInternalServerError(err error) *InternalServerError {
	return &InternalServerError{
		Err: err,
	}
}

func (e *InternalServerError) Error() string {
	return e.Err.Error()
}

func (e *InternalServerError) StatusCode() int {
	return 500
}

func (e *InternalServerError) Message() string {
	return "internal server error"
}

type BadRequestError struct {
	Err error
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		Err: errors.New(message),
	}
}

func (e *BadRequestError) Error() string {
	return e.Err.Error()
}

func (e *BadRequestError) StatusCode() int {
	return 400
}

func (e *BadRequestError) Message() string {
	return e.Err.Error()
}

// type NotFoundError struct {
// 	Err error
// }

// func (e *NotFoundError) Error() string {
// 	return e.Err.Error()
// }

// func (e *NotFoundError) StatusCode() int {
// 	return 404
// }

// func (e *NotFoundError) Message() string {
// 	return "Resource not found"
// }
