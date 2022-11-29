package config

import "errors"

var (
	ErrNotFound  = errors.New("not found")
	AlreadyExist = errors.New("already exitst")
	DublicateKey = errors.New("duplicate key value")
)
