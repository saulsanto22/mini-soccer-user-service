package error

import "errors"

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrSQLError            = errors.New("SQL Error")
	ErrTooManyRequest      = errors.New("Too Many Request")
	ErrUnauthorized        = errors.New("Unauthorized")
	ErrInvalidToken        = errors.New("Invalid Token")
	ErrForbidden           = errors.New("Forbidden")
	ErrBadRequest          = errors.New("Bad Request")
)

var General = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequest,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
	ErrBadRequest,
}
