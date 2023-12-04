// @Author huzejun 2023/12/4 23:39:00
package _case

import "fmt"

type Dog struct {
	animal
}

func NewDog() AnimalI {
	return &Dog{}
}

func (a Dog) Each() {
	fmt.Println("狗吃肉包子")
}

func (a Dog) Drink() {
	fmt.Println("狗喝白开水")
}

func (a Dog) Sleep() {
	fmt.Println("狗睡在窝里")
}

func (a Dog) Run() {
	fmt.Println("狗是4条腿跑")
}
