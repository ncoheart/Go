package main

import (
	"fmt"
	"reflect"
	"time"
)

// 将NewInt定义为int类型
type NewInt int

// 将int取一个别名叫IntAlias
type IntAlias = int

type MyDuration = time.Duration

// func (m MyDuration) EasySet(a string) {
// 非本地类不能定义方法
// }

type MynewDuration time.Duration

func (m MynewDuration) EasySet(a string) {
	// 这种可以
}

// 定义商标结构
type Brand struct {
}

// 为商标结构添加Show()方法
func (t Brand) Show() {

}

// 为Brand定义一个别名FakeBrand
type FakeBrand = Brand

// 定义车辆结构
type Vehicle struct {
	FakeBrand
	Brand
}

func main() {
	var a1 NewInt
	fmt.Printf("a type: %T\n", a1)

	var a2 IntAlias
	fmt.Printf("a2 type: %T\n", a2)

	var a Vehicle
	a.FakeBrand.Show()

	ta := reflect.TypeOf(a)

	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)
		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}
}
