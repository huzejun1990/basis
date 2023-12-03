// @Author huzejun 2023/11/19 19:45:00
package _case

import "fmt"

func VarDeclareCase() {
	// 通过var关键词声明变量
	var i, z, x int = 1, 2, 3
	// 通过var关键词声明变量，并赋值
	var j int = 100
	// 通过var关键词来赋值
	var f float32 = 100.123
	// 通过 := 以推断的方式定义变量并赋值
	b := true
	// 数组
	var arr = [5]int{1, 2, 3, 4, 5}
	arr1 := [...]int{2, 3, 4, 5, 6}
	var arr2 [5]int
	arr2[2] = 4
	arr2[3] = 5
	fmt.Println(i, z, x, j, f, b, arr, arr1, arr2)

	//指针类型，用于表示变量地址的类型
	var intPtr *int
	var floatPrt *float64
	var i1 = 100
	f1(&i1)

	//接口类型
	var inter interface{}
	inter = i1
	inter = f
	fmt.Println(intPtr, floatPrt, i1, inter)
}

func f1(i *int) {
	*i = *i + 1
}
