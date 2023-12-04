// @Author huzejun 2023/12/4 23:33:00
package _case

import "fmt"

type Cat struct {
	animal
}

func NewCat() AnimalI {
	return &Cat{}
}

func (c *Cat) Each() {
	fmt.Println("猫吃老鼠")
}
