// @Author huzejun 2023/11/30 2:49:00
package _case

import (
	"fmt"
	"strconv"
	"time"
	"unsafe"
)

func ConvertCase() {
	// 数字类型的转换

	// 不同类型的数据转换
	// 接口类型转其他类型
	// 数字类型的转换
	var num1 int = 100
	fmt.Println(int64(num1))
	var num2 int64 = 100
	fmt.Println(int(num2))

	// 字符串与数字类型转换
	var num3 = 100
	fmt.Println(strconv.Itoa(num3) + "abc")
	var str1 = "100"
	fmt.Println(strconv.Atoi(str1))

	var num4 int64 = 1010
	fmt.Println(strconv.FormatInt(num4, 10))

	var str2 = "1010"
	fmt.Println(strconv.ParseInt(str2, 10, 64))

	// 字符串与[]byte 转换
	var str3 = "今天天气很好"
	bytes1 := []byte(str3)
	fmt.Println(bytes1)

	fmt.Println(string(bytes1))

	// 字符串与rune转换
	// rune 为 int32类型
	// 将字符串转换为rune切片，实际上rune切片中存储了字符串的Unicode码点
	var rune1 = []rune(str3)
	fmt.Println(rune1)
	fmt.Println(string(rune1))
	fmt.Println(string(rune1[3]))
	fmt.Println([]int32(str3))

	// 接口类型转具体类型
	// 断言
	var inf interface{} = 100
	var infStruct interface{} = user{Name: "nick"}
	i, ok := inf.(int)
	fmt.Println(i, ok)
	u, ok := infStruct.(user)
	fmt.Println(u, ok)

	//时间类型转字符串
	var t time.Time
	t = time.Now()
	timeStr := t.Format("2006-01-02 15:04:05Z07:00")
	fmt.Println(timeStr)
	//字符串转时间
	t2, _ := time.Parse("2006-01-02 15:04:05Z07:00", timeStr)
	fmt.Println(t2)

	//uintptr
	u1 := user{}
	// unsafe.Pointer 是一个通用的指针类型，不能用于计算
	uPtr := unsafe.Pointer(&u1)
	namePtr := unsafe.Pointer(uintptr(uPtr) + unsafe.Offsetof(u1.Name))
	*(*string)(namePtr) = "nick"
	fmt.Println(u1)

}
