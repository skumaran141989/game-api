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

func (apierror *APIError) GetCode() int64 {
	return apierror.code
}

func (apierror *APIError) GetMessage() string {
	return apierror.message
}

func (apierror *APIError) GetException() error {
	return apierror.exception
}
