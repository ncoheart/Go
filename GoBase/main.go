package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 变量的声明
	var name int = 10
	name++
	fmt.Printf("%d\n", name)

	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17
	var damage = float32(attack-defence) * damageRate
	fmt.Println(damage)

	a, b := 1, 3
	a, c := 2, 4
	fmt.Printf("%d %d %d\n", a, b, c)

	fmt.Println(utf8.RuneCountInString("忍者"))
	fmt.Println(utf8.RuneCountInString("龙忍出鞘,fight!"))

	theme := "狙击 start"

	for _, s := range theme {
		fmt.Printf("Unicode: %c  %d\n", s, s)
	}

	var progress = 2
	var target = 8

	// 两参数格式化
	title := fmt.Sprintf("已采集%d个药草, 还需要%d个完成任务", progress, target)

	fmt.Println(title)

	pi := 3.14159
	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)

	fmt.Println(variant)

	// 匿名结构体声明, 并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}

	fmt.Printf("使用'%%+v' %+v\n", profile)

	fmt.Printf("使用'%%#v' %#v\n", profile)

	fmt.Printf("使用'%%T' %T\n", profile)
}
