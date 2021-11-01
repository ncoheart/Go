package main

import "fmt"

type People struct {
	name  string
	child *People
}

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

func main() {
	// 使用“键值对”初始化结构体
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	fmt.Println(relation.child.child.name)

	// 使用多个值的列表初始化结构体
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}
	fmt.Println(addr)
}
