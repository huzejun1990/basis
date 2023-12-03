// @Author huzejun 2023/11/2 3:55:00
package _case

import "fmt"

func SimpleCase() {
	var a, b = 3, 4
	var c, d float64 = 5, 6
	fmt.Println("不使用泛型，数字比较：", getMaxNumInt(a, b))
	fmt.Println("不使用泛型，数字比较：", getMaxNumFloat(c, d))

	//由编译器推断输入的类型
	//fmt.Println("使用泛型，数字比较：",getMxNum[int](a,b))
	fmt.Println("使用泛型，数字比较：", getMaxNum(a, b))
	// 显示指定传入的类型
	fmt.Println("使用泛型，数字比较：", getMaxNum[float64](c, d))

}

func getMaxNumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxNumFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func getMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b

}

type CusNumT interface {
	// 支持 uint8、int32、float64与int64及其衍生类型
	// ~  表示支持类型的衍生类型
	// | 表示取并集
	// 多选之间取 交集
	uint8 | int32 | float64 | ~int64
	int32 | float64 | ~int64 | uint16
}

// MyInt64 为int64的衍生类型，是具有基础类型int64的新类型
type MyInt64 int64

// MyInt32 为 int32 的别名，与int32 是同一类型
type MyInt32 = int32

func CusNumTCase() {
	var a, b int32 = 3, 4
	var a1, b1 MyInt32 = a, b
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(a, b))
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(a1, b1))

	var c, d float64 = 5, 6
	//由编译器推断输入的类型
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(c, d))
	// 显示指定传入的类型
	fmt.Println("自定义泛型，数字比较：", getMaxNum[float64](c, d))

	var e, f int64 = 7, 8
	var g, h MyInt64 = 7, 8

	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(e, f))
	fmt.Println("自定义泛型，数字比较：", getMaxCusNum(g, h))
}

func getMaxCusNum[T CusNumT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// 内置类型
func BuiltInCase() {
	var a, b string = "abc", "efg"
	fmt.Println("内置 comparable  泛型类型约束", getBuiltInComparable(a, b))
	var c, d float64 = 100, 100
	fmt.Println("内置 comparable  泛型类型约束", getBuiltInComparable(c, d))

	var f = 100.123
	printBuildInAny(f)
	printBuildInAny(a)
}

func getBuiltInComparable[T comparable](a, b T) bool {
	// comparable 类型，只支持 == != 两个操作
	if a == b {
		return true
	}
	return false
}

func printBuildInAny[T any](a T) {
	fmt.Println("内置 any 泛型类型约束", a)
}
