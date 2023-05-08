package xerror

import (
	"strings"

	"github.com/Sunshine958/commonx/errors/xcode"
)

const (
	commaSeparatorSpace = ", "
)

// New creates and returns an error which is formatted from given text.
func New(text string) error {
	return &ErrorModel{
		text: text,
		code: xcode.CodeNil,
	}
}

// NewCode creates and returns an error that has error code and given text.
func NewCode(code xcode.Code, text ...string) error {
	return &ErrorModel{
		text: strings.Join(text, commaSeparatorSpace),
		code: code,
	}
}

// WrapCode wraps error with code and text.
// It returns nil if given err is nil.
func WrapCode(code xcode.Code, err error, text ...string) error {
	if err == nil {
		return nil
	}
	return &ErrorModel{
		error: err,
		text:  strings.Join(text, commaSeparatorSpace),
		code:  code,
	}
}

// Wrap wraps error with text. It returns nil if given err is nil.
// Note that it does not lose the error code of wrapped error, as it inherits the error code from it.
func Wrap(err error, text string) error {
	if err == nil {
		return nil
	}
	return &ErrorModel{
		error: err,
		text:  text,
		code:  Code(err),
	}
}

// Code returns the error code of current error.
// It returns `CodeNil` if it has no error code neither it does not implement interface Code.
func Code(err error) xcode.Code {
	if err == nil {
		return xcode.CodeNil
	}
	if e, ok := err.(*ErrorModel); ok {
		return e.Code()
	}
	return xcode.CodeNil
}
