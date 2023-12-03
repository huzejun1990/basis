// @Author huzejun 2023/11/20 15:08:00
package _case

import "fmt"

type user struct {
	Name string
	Age  uint
	Addr Address
}

type Address struct {
	Province string
	City     string
}

func StructCase() {
	// 值类型
	u := user{
		Name: "nick",
		Age:  18,
	}
	f2(u)
	fmt.Println(u)
	// 指针类型
	u1 := &user{
		Name: "nick1",
		Age:  19,
	}
	fmt.Println(u1)
	//指针类型
	u2 := new(user)
	u2.Name = "nick3"
	u2.Age = 20
	fmt.Println(u2, u2.Addr.Province)

	// 结构体为值类型，定义变量后将默认初始化
	var u3 user
	fmt.Println(u3)

}

func f2(u user) {
	u.Age = 28
}
