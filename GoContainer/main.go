package main

import "fmt"

func main() {
	var a [4]int = [4]int{1, 2}
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}

	q := [...]int{1, 2, 3, 4}
	fmt.Printf("%T\n", q)
	fmt.Println(a == q)

	var array [4][2]int
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println(array[1][1])

	// 切片
	fmt.Println(a[1:3])

	var strList []string
	fmt.Println(strList)

	var numListEmpty = []int{}
	fmt.Println(numListEmpty == nil) //false

	m1 := make([]int, 2)
	m2 := make([]int, 2, 10)
	fmt.Println(m1, m2)
	fmt.Println(len(m1), len(m2))

	var ap []int
	ap = append(ap, 1)
	ap = append(ap, 2, 3)
	ap = append(ap, []int{4, 5, 6, 7}...) //需要解包

	ap = append([]int{-3, -2, -1, 0}, ap...)
	fmt.Println(ap)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
}
