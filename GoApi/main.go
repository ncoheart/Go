package main

import "fmt"

func main() {
	var invoker Invoker

	s := new(Struct)

	invoker = s
	invoker.Call("hello")

	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})

	invoker.Call("hello")

	// 闭包修改引用变量
	str := "hello world"
	foo := func() {
		str = "hello dude"
		fmt.Println(str)
	}
	foo()
}

type Invoker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

type FuncCaller func(interface{})

func (f FuncCaller) Call(p interface{}) {
	f(p)
}
