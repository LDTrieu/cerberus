package errors

import (
	"fmt"

	"github.com/pkg/errors"
)

// List of errors
var (
	ErrRecordNotFound = errors.New("record not found")
)

const (
	//TxtErrHappennedTemplate ...
	TxtErrHappennedTemplate       = "Có lỗi xảy ra (%s)"
	TxtErrRequestNotFoundTemplate = "Yêu cầu không tồn tại (%s)"
)

// Wrap error
func Wrap(err error, format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	return errors.Wrap(err, message)
}

// New create a new error
func New(format string, args ...interface{}) error {
	message := fmt.Sprintf(format, args...)
	return errors.New(message)
}

// NewInternalWithCode ...
func NewInternalWithCode(code string) error {
	return fmt.Errorf(TxtErrHappennedTemplate, code)
}

// NewRequestNotFoundWithCode ...
func NewRequestNotFoundWithCode(code string) error {
	return fmt.Errorf(TxtErrRequestNotFoundTemplate, code)
}

// Cause returns the underlying cause of the error, if possible.
func Cause(err error) error {
	return errors.Cause(err)
}

// IsRecordNotFoundError ...
func IsRecordNotFoundError(err error) bool {
	if err == ErrRecordNotFound {
		return true
	}
	if cause := errors.Cause(err); cause == ErrRecordNotFound {
		return true
	}

	return false
}
