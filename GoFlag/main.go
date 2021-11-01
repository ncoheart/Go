package main

import (
	"flag"
	"fmt"
)

var mode = flag.String("mode", "", "process mode")

func main() {
	flag.Parse()

	fmt.Println(*mode)

	str := new(string)
	*str = "Go语言教程"
	fmt.Println(*str)
}
