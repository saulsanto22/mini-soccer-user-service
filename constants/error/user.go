package error

import "errors"

var (
	ErrUserNotFound      = errors.New("User Not Found")
	ErrPasswordInCorrect = errors.New("Password Incorrect")
	ErrUsernameExist     = errors.New("Username Already Exist")
	ErrEmailExist        = errors.New("Email Already Exist")
	ErrPasswordDontMatch = errors.New("Password Don't Match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrPasswordInCorrect,
	ErrUsernameExist,
	ErrPasswordDontMatch,
}
