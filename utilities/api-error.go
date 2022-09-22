package utilities

type APIError struct {
	code      int64
	message   string
	exception error
}

func NewAPIError(code int64, message string, exception error) *APIError {
	return &APIError{
		code:      code,
		message:   message,
		exception: exception,
	}
}
