package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("fist")
	l.PushBack(67)
	element := l.PushBack(56)
	l.PushBack(46)
	l.PushBack(34)
	l.PushBack(89)

	// 在fist之后添加high
	l.InsertAfter("high", element)
	// 在fist之前添加noon
	l.InsertBefore("noon", element)

	l.Remove(element)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
