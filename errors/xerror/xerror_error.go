package xerror

import "github.com/Sunshine958/commonx/errors/xcode"

type ErrorModel struct {
	error error // Wrapped error.
	text  string
	code  xcode.Code
}

func (err *ErrorModel) Error() string {
	if err == nil {
		return ""
	}
	errStr := err.text
	if errStr == "" && err.code != nil {
		errStr = err.code.Message()
	}
	if err.error != nil {
		if errStr != "" {
			errStr += ": "
		}
		errStr += err.error.Error()
	}
	return errStr
}

func (err *ErrorModel) Code() xcode.Code {
	if err == nil {
		return xcode.CodeNil
	}
	return err.code
}
