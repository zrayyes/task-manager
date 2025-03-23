package repositories

import "errors"

var ErrNotFound error = errors.New("not found")

var ErrAlreadyExists error = errors.New("already exists")
