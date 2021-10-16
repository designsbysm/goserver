package database

import (
	"errors"
)

var ErrPasswordRequired = errors.New("password required for user")
