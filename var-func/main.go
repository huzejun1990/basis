// @Author huzejun 2023/11/1 21:36:00
package main

import (
	_case "basis/var-func/case"
	"fmt"
)

func main() {
	a := 10
	b := 20
	fmt.Println(_case.SumCase(a, b))
	fmt.Println(a, b)
	_case.ReferenceCase(a, &b)
	fmt.Println(a, b)

	fmt.Println(_case.G)
	_case.ScopeCase(a, b)
	fmt.Println(_case.G)

	user := _case.NewUser("nick", 18)
	fmt.Println(user.GetName(), user.GetAge())
}
