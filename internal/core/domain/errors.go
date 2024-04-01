package domain

import (
	"errors"
)

var ErrNoRecord = errors.New("No records were found")

var ErrInvalidCredentials = errors.New("Invalid credentials")
