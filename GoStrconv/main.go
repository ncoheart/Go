package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("a" + strconv.Itoa((32)))

	i, _ := strconv.Atoi("33")
	fmt.Println(3 + i)
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println("转换失败！")
	}

	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.14.15", 64)
	in, err := strconv.ParseInt("-42", 10, 64) // 以十进制解析，保存为int64
	fmt.Println(b, f, in)

	s := strconv.FormatBool(true)
	println(s)
	s = strconv.FormatFloat(3.1414, 'E', -1, 64)
	println(s)
	s = strconv.FormatInt(-42, 16)
	println(s)

	b10 := []byte("int (base 10):")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))
	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
}
