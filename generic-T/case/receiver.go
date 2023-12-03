// @Author huzejun 2023/11/7 2:51:00
package _case

import "fmt"

// 泛型结构体
type MyStruct[T interface{ *int | *string }] struct {
	Name string
	Data T
}

// 泛型 receiver
func (myStruct MyStruct[T]) GetData() T {
	return myStruct.Data
}

func ReceiverCase() {
	data := 18
	myStruct := MyStruct[*int]{
		Name: "nick",
		Data: &data,
	}
	data1 := myStruct.GetData()
	fmt.Println(*data1)

	str := "abcdefg"
	myStruct1 := MyStruct[*string]{
		Name: "nick",
		Data: &str,
	}
	str1 := myStruct1.GetData()
	fmt.Println(*str1)

}
