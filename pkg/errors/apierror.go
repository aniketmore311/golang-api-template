package errors

type APIError struct {
	status int
	code   ErrorCode
	detail string
	error
}

func (ae *APIError) GetStatus() int     { return ae.status }
func (ae *APIError) GetCode() ErrorCode { return ae.code }
func (ae *APIError) GetDetail() string  { return ae.detail }

func (ae *APIError) Error() string {
	return ae.error.Error()
}

func NewAPIError(err error, status int, code ErrorCode, detail string) *APIError {
	return &APIError{
		status,
		code,
		detail,
		err,
	}
}

var _ error = (*APIError)(nil)
