package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

// 根据键读取值
func readValue(key string) int {
	valueByKeyGuard.Lock()
	defer valueByKeyGuard.Unlock()

	return valueByKey[key]
}

func main() {
	fmt.Println("defer begin")

	defer fmt.Println(1)

	defer fmt.Println(2)

	defer fmt.Println(3)

	fmt.Println("defer end")

	// defer可以在函数退出时释放资源

	// 使用延迟释放文件句柄

}

func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer f.Close()
	// 取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时发生错误, 关闭文件并返回文件大小为0
	if err != nil {
		f.Close()
		return 0
	}
	// 取文件大小
	size := info.Size()

	// 返回文件大小
	return size
}
