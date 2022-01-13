package utils

import "os"

func IsFileExist(f string) bool {
	s, err := os.Stat(f)
	if err != nil {
		if os.IsExist(err) && !s.IsDir() {
			return true
		}
		return false
	}
	if !s.IsDir() {
		return true
	}
	return false
}
