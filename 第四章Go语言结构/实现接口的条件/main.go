package main

import (
	"fmt"
)

// 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error

	// 能否写入
	// CanWrite() bool
}

// 定义文件结构，用于实现DataWriter
type file struct {
}

// 实现DataWriter接口的WriteData方法
func (d *file) WriteData(data interface{}) error {
	fmt.Println("WriteData:", data)
	return nil
}

func main() {
	// 条件1：接口的方法与实现接口的类型方法格式一致
	// 实例化file
	f := new(file)
	var writer DataWriter
	writer = f
	writer.WriteData("data22")

	// 条件2：接口中所有方法均被实现

}
