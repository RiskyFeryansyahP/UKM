package utils

// Error model customize error
type Error struct {
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
}

// WrapErrorJson is return message error and status code
func WrapErrorJson(err error, status int) *Error {
	return &Error{
		StatusCode: status,
		Message:    err.Error(),
	}
}
