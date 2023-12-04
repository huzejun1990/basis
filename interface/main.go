// @Author huzejun 2023/12/4 20:58:00
package main

import _case "basis/interface/case"

func main() {
	cat := _case.NewCat()
	animalLife(cat)
	dog := _case.NewDog()
	animalLife(dog)
	dove := _case.NewDove()
	animalLife(dove)
}

func animalLife(a _case.AnimalI) {
	a.Each()
	a.Drink()
	a.Sleep()
	a.Run()
}
