package errs

import "errors"

var (
	ErrUserAlreadyExists           = errors.New("user already exists")
	ErrNotFound                    = errors.New("not found")
	ErrIncorrectUsernameOrPassword = errors.New("incorrect username or password")
)
