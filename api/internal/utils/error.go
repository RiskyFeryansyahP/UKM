package utils

type Error struct {
	StatusCode int `json:"statuscode"`
	Message string `json:"message"`
}

func WrapErrorJson(err error, status int) *Error {
	return &Error{
		StatusCode: status,
		Message:    err.Error(),
	}
}