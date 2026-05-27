package apperror

type AppError struct {
	message   string
	errorType ErrorType
}

func NewAppError(message string, errorType ErrorType) *AppError {
	return &AppError{
		message:   message,
		errorType: errorType,
	}
}

func (a *AppError) Error() string {
	return a.message
}

func (a *AppError) GetErrorType() ErrorType {
	return a.errorType
}
