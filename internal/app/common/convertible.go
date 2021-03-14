package common

type Convertible interface {
	// 返回要上报的 JSON 的 byte 数组，如果发生错误，log.Error() 并返回 nil 即可，无需 panic() 或 log.Fatal()
	Bytes() []byte
}
