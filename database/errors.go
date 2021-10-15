package database

import (
	"errors"
)

var errorPasswordRequired = errors.New("User error: password is required")
var errorPropertyRequired = errors.New("Setting error: property is required")
