package errors

import (
	"errors"
	"io"
	"testing"
)

func TestErrorChain(t *testing.T) {
	err := Top()
	if err == nil {
		t.Fatal("expect error, got nil")
	}

	// 逐层 unwrap
	expectOps := []Op{"top", "middle", "root"}
	for i, op := range expectOps {
		if e, ok := err.(*MyError); !ok {
			t.Fatalf("layer %d: not MyError", i)
		} else if e.Op != op {
			t.Fatalf("layer %d: got op %s, want %s", i, e.Op, op)
		}
		err = errors.Unwrap(err)
	}
	// 最后一层是 io.EOF
	if !errors.Is(err, io.EOF) {
		t.Fatalf("bottom error should be io.EOF, got %v", err)
	}
}

func TestMyError_Unwrap(t *testing.T) {
	e := &MyError{Op: "a", Err: io.EOF}
	if errors.Unwrap(e) != io.EOF {
		t.Fail()
	}
}