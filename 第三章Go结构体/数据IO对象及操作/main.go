package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("C语言中文网,Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)

	// Read()方法
	// var buf [128]byte
	// n, err := r.Read(buf[:])
	// fmt.Println(string(buf[:n]), n, err)
	// 前面读完之后指针会指向最后，此时如果再进行读操作会报错，需要进行注释

	// ReadByte()方法
	c, err := r.ReadByte()
	fmt.Println(string(c), err)

	// ReadBytes()方法
	var delim byte = ',' // 用于指定分割字符
	line, err := r.ReadBytes(delim)
	fmt.Println(string(line), err)

	// ReadLine()方法，一般使用ReadBytes('\n')或者ReadString('\n')替代
	data2 := []byte("Golang is a beautiful language. \r\n I like it!")
	rd = bytes.NewReader(data2)
	r = bufio.NewReader(rd)
	line, prefix, err := r.ReadLine()
	fmt.Println(string(line), prefix, err)

	// ReadRune()方法

	// ReadSlice()方法
	data3 := []byte("C语言中文网,Go语言入门级教程")
	rd = bytes.NewReader(data3)
	r = bufio.NewReader(rd)
	delim = ','
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
}
