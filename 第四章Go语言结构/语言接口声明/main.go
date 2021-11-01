package main

type writer interface {
	// 方法名(参数列表) 返回值列表
	Write([]byte) error
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func main() {

}
