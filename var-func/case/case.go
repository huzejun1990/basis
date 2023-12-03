// @Author huzejun 2023/11/1 21:37:00
package _case

import "errors"

//形参
// 局部变量
// 全局变量

func SumCase(a, b int) (sum int, err error) {
	if a <= 0 && b >= 0 {
		err = errors.New("两数相加不能同时小于0")
		return
	}
	sum = a + b
	return
}

// 值传递与引用传递
func ReferenceCase(a int, b *int) {
	a += 1
	*b += 1
}

// 全局变量
var g int

var G int

// 变量作用域
func ScopeCase(a, b int) {
	c := 100
	g = a + b + c
	G = g

}

type User struct {
	Name string
	Age  uint
}

func NewUser(name string, age uint) *User {
	return &User{Name: name, Age: age}
}

// 获取user name属性
func (u *User) GetName() string {
	return u.Name
}

// 获取user age属性
func (u *User) GetAge() uint {
	return u.Age
}
