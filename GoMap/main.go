package main

import (
	"fmt"
	"sync"
)

func main() {
	var mapList map[string]int
	mapList = map[string]int{"one": 1, "two": 2}
	fmt.Printf("Map literal at \"one\" is: %d\n", mapList["one"])

	var scene sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	fmt.Println(scene.Load("london"))

	scene.Delete("london")

	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
