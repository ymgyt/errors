package errors

import (
	"bytes"
	"fmt"
)

// ErrCode -
type ErrCode int

//go:generate stringer -type ErrCode error.go
const (
	OK ErrCode = iota
	Invalid
	Conflict
	Timeout
	Internal
	External
	NotFound
	Unauthorized
	Unauthenticated
	RateLimit
	Undefined
)

const (
	defaultMessage = "an internal error has occured"
)

// E -
type E struct {
	Code    ErrCode
	Message string
	Op      string
	Err     error
}

// Code -
func Code(err error) ErrCode {
	if err == nil {
		return OK
	} else if e, ok := err.(*E); ok && e.Code != OK {
		return e.Code
	} else if ok && e.Err != nil {
		return Code(e.Err)
	}
	return Undefined
}

// Message -
func Message(err error) string {
	if err == nil {
		return ""
	} else if e, ok := err.(*E); ok && e.Message != "" {
		return e.Message
	} else if ok && e.Err != nil {
		return Message(e.Err)
	}
	return defaultMessage
}

// Error -
func (e *E) Error() string {
	var b bytes.Buffer

	if e.Op != "" {
		fmt.Fprintf(&b, "%s: ", e.Op) //nolint: gas
	}

	if e.Err != nil {
		b.WriteString(e.Err.Error())
	} else {
		if e.Code != OK {
			fmt.Fprintf(&b, "<%s> ", e.Code) //nolint: gas
		}
		b.WriteString(e.Message)
	}
	return b.String()
}
