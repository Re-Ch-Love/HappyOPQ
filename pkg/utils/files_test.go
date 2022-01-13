package utils

import (
	"os"
	"testing"
)

func TestIsExist(t *testing.T) {
	a := Asserter{T: t}

	err := os.Mkdir("a", os.ModePerm)
	a.AssertEqual(err, nil, err)
	a.AssertFalse(IsFileExist("a"), "")
	a.AssertFalse(IsFileExist("./a"), "")
	err = os.Remove("a")
	a.AssertEqual(err, nil, err)

	f, err := os.Create("a")
	a.AssertEqual(err, nil, err)
	a.AssertTrue(IsFileExist("a"), "")
	a.AssertTrue(IsFileExist("./a"), "")
	err = f.Close()
	a.AssertEqual(err, nil, err)
	err = os.Remove("a")
	a.AssertEqual(err, nil, err)
}
