package common

type Convertible interface {
	// Convert 返回QQ号和转换后的结构体
	Convert() (int64, interface{})
}
