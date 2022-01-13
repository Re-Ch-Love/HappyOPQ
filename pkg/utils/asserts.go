package utils

import (
	"runtime/debug"
	"testing"
)

type Asserter struct {
	T *testing.T
}

func (a *Asserter) AssertTrue(v bool, msg ...interface{}) {
	if !v {
		a.T.Fatal(append([]interface{}{string(debug.Stack()), "\n"}, msg...)...)
	}
}

func (a *Asserter) AssertFalse(v bool, msg ...interface{}) {
	if v {
		a.T.Fatal(append([]interface{}{string(debug.Stack()), "\n"}, msg...)...)
	}
}

func (a *Asserter) AssertEqual(v1 interface{}, v2 interface{}, msg ...interface{}) {
	if !(v1 == v2) {
		a.T.Fatal(append([]interface{}{string(debug.Stack()), "\n"}, msg...)...)
	}
}

func (a *Asserter) AssertNotEqual(v1 interface{}, v2 interface{}, msg ...interface{}) {
	if v1 == v2 {
		a.T.Fatal(append([]interface{}{string(debug.Stack()), "\n"}, msg...)...)
	}
}
