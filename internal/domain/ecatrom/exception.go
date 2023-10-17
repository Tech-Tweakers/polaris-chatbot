package ecatrom

import (
	"errors"
	"fmt"
)

type DomainError struct {
	StatusCode int
	Err        error
	Retryable  bool
	Message    string
}

type domainError struct {
	StatusCode  int
	Description string
	Retryable   bool
}

var NotModified *domainError = &domainError{StatusCode: 304, Description: "not modified", Retryable: false}
var BadRequest *domainError = &domainError{StatusCode: 400, Description: "bad request", Retryable: false}
var NotFound *domainError = &domainError{StatusCode: 404, Description: "not found", Retryable: false}

func (r *DomainError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func DomainErrorFactory(err *domainError, message string) error {
	return &DomainError{
		StatusCode: err.StatusCode,
		Message:    message,
		Err:        errors.New(err.Description),
	}
}
