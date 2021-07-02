package errors

import "errors"

type ErrorType int

const (
	ServerErr ErrorType = iota
	ClientErr
)

type Error struct {
	Type ErrorType
	Err  error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func New(errType ErrorType, text string, err error) error {
	e := &Error{
		Type: errType,
	}
	if len(text) != 0 {
		e.Err = errors.New(text)
	} else if err != nil {
		e.Err = err
	}
	return e
}
