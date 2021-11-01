package main

import "fmt"

func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}
func main() {
	str := "hello world"
	foo := func() {
		str = "hello dude"
	}
	foo()
	print(str)

	// 闭包的记忆效应
	// 创建一个累加器, 初始值为1
	accumulator := Accumulate(1)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)
	// 创建一个累加器, 初始值为1
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)
}
