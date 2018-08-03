package errors

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {
	tcs := []struct {
		name string
		e    error
		want ErrCode
	}{
		{"nil is OK", nil, OK},
		{"simple", &E{Code: Invalid}, Invalid},
		{"nest", &E{Err: &E{Code: Conflict}}, Conflict},
		{"override", &E{Code: Internal, Err: &E{Code: Conflict}}, Internal},
		{"undefined", fmt.Errorf(""), Undefined},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got, want := Code(tc.e), tc.want; got != want {
				t.Errorf("got %s, want %s", got, want)
			}
		})
	}
}

func TestMessage(t *testing.T) {
	tcs := []struct {
		name string
		e    error
		want string
	}{
		{"nil is empty", nil, ""},
		{"simple", &E{Message: "msg"}, "msg"},
		{"nest", &E{Err: &E{Message: "nested msg"}}, "nested msg"},
		{"override", &E{Message: "override", Err: &E{Message: "never"}}, "override"},
		{"default", fmt.Errorf(""), defaultMessage},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got, want := Message(tc.e), tc.want; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

func TestE_Error(t *testing.T) {
	tcs := []struct {
		name string
		e    error
		want string
	}{
		{"simple message", &E{Message: "msg"}, "msg"},
		{"simple code and message", &E{Code: Invalid, Message: "msg"}, fmt.Sprintf("<%s> %s", Invalid.String(), "msg")},
		{"operation", &E{Op: "find users", Code: Invalid, Message: "invalid user id format"}, fmt.Sprintf("%s: <%s> %s", "find users", Invalid.String(), "invalid user id format")},
		{"nest", &E{Op: "create_user", Err: &E{Op: "check_input", Err: &E{Code: Invalid, Message: "invalid user id"}}}, fmt.Sprintf("%s: %s: <%s> %s", "create_user", "check_input", Invalid.String(), "invalid user id")},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			if got, want := tc.e.Error(), tc.want; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}
