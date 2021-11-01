package main

type Point struct {
	X int
	Y int
}

func main() {
	var p Point
	p.X = 10
	p.Y = 20

	//创建指针类型的结构体
	tank := new(Point)
	tank.X = 40
	tank.Y = 50

	// 取结构体的地址实例化
	cmd := &Point{}
	cmd.X = 70
	cmd.Y = 100
}
