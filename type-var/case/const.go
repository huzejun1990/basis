// @Author huzejun 2023/11/20 14:32:00
package _case

import (
	"fmt"
)

const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
	TB
)

type Gender uint

const (
	FEMALE Gender = iota
	MALE
	THIRD
)

func ConstAndEnumCase() {
	const (
		A = 10
		B = 20
	)
	size := MB
	var gender Gender = MALE
	fmt.Println(A, B, gender, size)
}
