// Code generated by "stringer -type ErrCode error.go"; DO NOT EDIT.

package errors

import "strconv"

const _ErrCode_name = "OKInvalidConflictTimeoutInternalExternalNotFoundUnauthorizedUnauthenticatedRateLimitUndefined"

var _ErrCode_index = [...]uint8{0, 2, 9, 17, 24, 32, 40, 48, 60, 75, 84, 93}

func (i ErrCode) String() string {
	if i < 0 || i >= ErrCode(len(_ErrCode_index)-1) {
		return "ErrCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrCode_name[_ErrCode_index[i]:_ErrCode_index[i+1]]
}