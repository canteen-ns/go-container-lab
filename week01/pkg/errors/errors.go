package errors

import (
	"fmt"
	"io"
)

type Op string

type MyError struct {
	Op  Op
	Msg string
	Err error
}

func (e *MyError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s %s\n\tcaused by: %v", e.Op, e.Msg, e.Err)
	}

	return fmt.Sprintf("%s %s", e.Op, e.Msg)
}

func (e *MyError) Unwrap() error {
	return e.Err
}

func root() error {
	e := &MyError{Op: "root", Msg: "io.EOF", Err: io.EOF}

	return e
}

func middle() error {
	if err := root(); err != nil {
		return &MyError{Op: "middle", Msg: "call root fail", Err: err}
	}

	return nil
}

func Top() error {
	if err := middle(); err != nil {
		return &MyError{Op: "top", Msg: "call middle fail", Err: err}
	}

	return nil
}
