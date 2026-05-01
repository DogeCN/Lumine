package errors

import (
	"errors"
	"unsafe"
)

var New = errors.New

type opError struct {
	op  string
	err error
}

func (e *opError) Error() string {
	return e.op + ": " + e.err.Error()
}

func (e *opError) Unwrap() error {
	return e.err
}

func WithStr(op string, err error) error {
	return &opError{
		op:  op,
		err: err,
	}
}

// Modified from errors.joinError
type joinError struct {
	errs []error
}

func (e *joinError) Error() string {
	if len(e.errs) == 1 {
		return e.errs[0].Error()
	}
	b := []byte(e.errs[0].Error())
	for _, err := range e.errs[1:] {
		b = append(b, ';', ' ')
		b = append(b, err.Error()...)
	}
	return unsafe.String(&b[0], len(b))
}

func (e *joinError) Unwrap() []error {
	return e.errs
}

func Join(errs ...error) error {
	n := 0
	for _, err := range errs {
		if err != nil {
			n++
		}
	}
	if n == 0 {
		return nil
	}
	e := &joinError{
		errs: make([]error, 0, n),
	}
	for _, err := range errs {
		if err != nil {
			e.errs = append(e.errs, err)
		}
	}
	return e
}