package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(hypot(3, 4))
	a, b := typedTwoValues()
	fmt.Println(a, b)
	a, b = namedRetValues()
	a, b = namedRetValuesNew()

	var f func()
	f = fire
	f()

	// Go语言字符串的链式处理——操作与数据分离的设计技巧
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}

	chain := []func(string) string{
		RemovePrefix,
		strings.TrimSpace,
		strings.ToUpper,
	}

	StringProccess(list, chain)

	for _, str := range list {
		fmt.Println(str)
	}

	// 匿名函数，并且可以通过()直接调用
	func(data int) {
		fmt.Println("hello", data)
	}(100)

	f1 := func(data int) {
		fmt.Println("hello", data)
	}
	f1(200)

	// 使用匿名函数打印切片内容
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

// 同一类型返回值
func typedTwoValues() (int, int) {
	return 1, 2
}

// 带有变量名的返回值
func namedRetValues() (a, b int) {
	a = 1
	b = 2
	return
}

// 也可写作
func namedRetValuesNew() (a, b int) {
	a = 1
	return a, 2
}

// 把函数作为值保存到变量中
func fire() {
	fmt.Println("fire")
}

func StringProccess(list []string, chain []func(string) string) {
	for index, str := range list {
		result := str

		for _, proc := range chain {
			result = proc(result)
		}

		list[index] = result
	}
}

func RemovePrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

// 遍历切片的每个元素, 通过给定函数进行元素访问
func visit(list []int, f func(int)) {

	for _, v := range list {
		f(v)
	}
}
