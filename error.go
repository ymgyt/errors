package errors

import (
	"bytes"
	"fmt"
)

// ErrCode -
type ErrCode int

const (
	OK ErrCode = iota
	INVALID
	CONFLICT
	TIMEOUT
	INTERNAL
	EXTERNAL
	NOTFOUND
	UNAUTHORIZED
	UNAUTHENTICATED
	RATELIMIT
	UNDEFINED
	MAXCODE
)

// String -
func (c ErrCode) String() string {
	switch c {
	case OK:
		return ""
	case INVALID:
		return "invalid"
	case CONFLICT:
		return "conflict"
	case TIMEOUT:
		return "timeout"
	case INTERNAL:
		return "internal"
	case EXTERNAL:
		return "external"
	case NOTFOUND:
		return "not_found"
	case UNAUTHORIZED:
		return "unauthorized"
	case UNAUTHENTICATED:
		return "unauthenticated"
	case RATELIMIT:
		return "rate_limit"
	case UNDEFINED:
		return "undefined"
	}
	return ""
}

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
	return UNDEFINED
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
	return "an internal error has occured."
}

// Error -
func (e *E) Error() string {
	var b bytes.Buffer

	if e.Op != "" {
		fmt.Fprintf(&b, "%s: ", e.Op)
	}

	if e.Err != nil {
		b.WriteString(e.Err.Error())
	} else {
		if e.Code != OK {
			fmt.Fprintf(&b, "<%s> ", e.Code)
		}
		b.WriteString(e.Message)
	}
	return b.String()
}
