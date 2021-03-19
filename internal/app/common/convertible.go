package common

type Convertible interface {
	// 返回 QQ 号和转换后的结构体
	Convert() (int64, interface{})
}
