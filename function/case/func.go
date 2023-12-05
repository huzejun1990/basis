// @Author huzejun 2023/12/6 2:52:00
package _case

import (
	"errors"
	"log"
)

func Sum(a, b int) (sum int, err error) {
	if a <= 0 && b <= 0 {
		err = errors.New("两数相加不同同时小于0")
		return 0, err
	}
	sum = a + b
	return
}

// 将函数作为类型
type SumFunc func(a, b int) (int, error)

// 将函数作为输入输出实现中间件
func LogMinddleware(in SumFunc) SumFunc {
	//返回的函数为闭包函数，其中in为闭包函数使用的外部函数内部变量
	return func(a, b int) (int, error) {
		log.Printf("日志中间件，记录操作数：a: %d, b: %d", a, b)
		return in(a, b)
	}
}

// 声明receiver为函数类型的方法，即函数类型的对象的方法
func (sum SumFunc) Accumulation(list ...int) (int, error) {
	s := 0
	var err error
	for _, i := range list {
		s, err = sum(s, i)
		if err != nil {
			return s, err
		}
	}
	return s, err
}
