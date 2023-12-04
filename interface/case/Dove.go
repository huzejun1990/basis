// @Author huzejun 2023/12/5 0:00:00
package _case

import "fmt"

type Dove struct {
}

func NewDove() AnimalI {
	return &Dove{}
}

func (d *Dove) Each() {
	fmt.Println("鸽子吃虫子")
}

func (d *Dove) Drink() {
	fmt.Println("鸽子喝水")
}

func (d *Dove) Sleep() {
	fmt.Println("鸽子睡树上")
}

func (d *Dove) Run() {
	fmt.Println("鸽子用翅膀飞")
}
