package apperror

type AppError interface {
	error
	StatusCode() int
	Message() string
}
