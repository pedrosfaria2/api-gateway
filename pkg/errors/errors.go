package errors

import (
	"fmt"
)

type Type string

const (
	ErrorTypeInternal      Type = "internal"
	ErrorTypeConfiguration Type = "configuration"
	ErrorTypeValidation    Type = "validation"
	ErrorTypeBackend       Type = "backend"
)

type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewError(t Type, message string, err error) *Error {
	return &Error{
		Type:    t,
		Message: message,
		Err:     err,
	}
}

func IsType(err error, t Type) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}
	return e.Type == t
}
