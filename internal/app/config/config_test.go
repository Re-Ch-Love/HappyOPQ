package config

import (
	"HappyOPQ/pkg/tests"
	"os"
	"testing"
)

func TestLoadConfigCustom(t *testing.T) {
	a := tests.Asserter{T: t}
	path := ".temp.yml"
	config := `
OPQBot:
    Port: 1234		
`
	f, err := os.Create(path)
	a.AssertEqual(err, nil, err)
	_, err = f.WriteString(config)
	a.AssertEqual(err, nil, err)
	err = f.Close()
	a.AssertEqual(err, nil, err)
	defer func() {
		err = os.Remove(path)
		a.AssertEqual(err, nil, err)
	}()

	if LoadConfig(path).OPQBot.Port != 1234 {
		t.FailNow()
	}
}

func TestLoadConfigDefault(t *testing.T) {
	if LoadConfig("").OPQBot.Port != 8080 {
		t.FailNow()
	}
}
