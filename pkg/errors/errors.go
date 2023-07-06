package errors

import "errors"

var (
	ErrNoSourcePath = errors.New("db: source path not provided")
	ErrNoSourceType = errors.New("db: source type not provided")
)