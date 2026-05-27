package apperror

type ErrorType int

const (
	UnExpected ErrorType = iota + 1
	Forbidden
	BadRequest
	NotFound
	Duplicate
)
