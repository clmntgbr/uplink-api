package domain

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrMaxProjectsReached    = errors.New("maximum number of projects reached")
	ErrProjectNotFound       = errors.New("project not found")
	ErrActiveProjectNotFound = errors.New("active project not found")
	ErrUserNotAuthenticated  = errors.New("user not authenticated")
)
