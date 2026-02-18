package errors

import "errors"

var (
	ErrUserNotFound            = errors.New("user not found")
	ErrProjectsNotFound        = errors.New("projects not found")
	ErrMaxProjectsReached      = errors.New("maximum number of projects reached")
	ErrProjectNotFound         = errors.New("project not found")
	ErrWorkflowsNotFound       = errors.New("workflows not found")
	ErrEndpointsNotFound       = errors.New("endpoints not found")
	ErrActiveProjectNotFound   = errors.New("active project not found")
	ErrUserNotAuthenticated    = errors.New("user not authenticated")
	ErrInvalidCredentials      = errors.New("invalid credentials")
	ErrInvalidQueryParams      = errors.New("invalid query params")
	ErrInvalidRequestBody      = errors.New("invalid request body")
	ErrInvalidTokenClaims      = errors.New("invalid token claims")
	ErrInvalidProjectID        = errors.New("invalid project id")
	ErrUserAlreadyExists       = errors.New("user already exists")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrInvalidToken            = errors.New("invalid token")
	ErrValidationFailed        = errors.New("validation failed")
)
