// @Author huzejun 2023/12/4 20:59:00
package _case

import "fmt"

// 声明 AnimalI 接口
// 定义 AnimalI 行为
type AnimalI interface {
	//吃
	Each()
	//喝
	Drink()
	//睡觉
	Sleep()
	//大奔跑
	Run()
}

type animal struct {
}

func (a animal) Each() {
	fmt.Println("Animal Each 接口默认实现")
}

func (a animal) Drink() {
	fmt.Println("Animal Drink 接口默认实现")
}

func (a animal) Sleep() {
	fmt.Println("Animal Sleep 接口默认实现")
}

func (a animal) Run() {
	fmt.Println("Animal Run 接口默认实现")
}
