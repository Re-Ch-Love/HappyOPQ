package tests

import (
	"testing"
)

func TestAsserter_AssertEqual(t *testing.T) {
	asserter := Asserter{T: t}
	asserter.AssertEqual(1, 1, "")
}

func TestAsserter_AssertNotEqual(t *testing.T) {
	asserter := Asserter{T: t}
	asserter.AssertNotEqual(1, 2, "")
}
