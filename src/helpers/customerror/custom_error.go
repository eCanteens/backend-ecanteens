package customerror

type CustomError struct {
	Message    string
	StatusCode int
}

// Error implements the error interface for CustomError
func (e *CustomError) Error() string {
	return e.Message
}

// New creates a new CustomError with a given message and status code
func New(message string, statusCode int) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: statusCode,
	}
}
