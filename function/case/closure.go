// @Author huzejun 2023/12/6 4:35:00
package _case

import (
	"fmt"
	"log"
)

func Fib(n int) int {
	if n <= 2 {
		log.Fatal("请选择大于2的位置")
	}

	t := tool()
	var res int
	for i := 0; i < n-2; i++ {
		res = t()
	}
	return res
}

// 斐波那契数列，X0+X1=X2...求n
func tool() func() int {
	var x0 = 0
	var x1 = 1
	var x2 = 0
	return func() int {
		x2 = x0 + x1
		x0 = x1
		x1 = x2
		return x2
	}
}

func ClosureTrap() {
	/* 错误的方式
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	*/

	for i := 0; i < 10; i++ {
		go func(j int) {
			fmt.Println(j)
		}(i)
	}
}
