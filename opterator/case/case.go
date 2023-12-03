// @Author huzejun 2023/12/2 22:30:00
package _case

import "fmt"

/*
// 运算符优先级
1. 一元运算符优先级最高. 例如：++ -- ！ ^
2. 二元运算符优先级从高到低
*		/		%		<<		>>		&		&^
+		-		|		^
==		!=		<		<=		>		>=
&&
||

*/

// 算术运算符
func ArithmeticCase() {
	var a = 21
	var b = 10
	var c int

	c = a + b
	fmt.Printf("a + b 的值为 %d \n", c)
	c = a - b
	fmt.Printf("a - b 的值为 %d \n", c)
	c = a * b
	fmt.Printf("a * b 的值为 %d \n", c)
	c = a / b
	fmt.Printf("a / b 的值为 %d \n", c)
	c = a % b
	fmt.Printf("a %% b 的值为 %d \n", c)
	a++
	fmt.Printf("a++ 的值为 %d \n", a)
	a--
	fmt.Printf("a-- 的值为 %d \n", a)

}

// 关系运算符
func RelationCase() {
	var a = 21
	var b = 10
	fmt.Println("a == b", a == b)
	fmt.Println("a > b", a > b)
	fmt.Println("a < b", a < b)
	fmt.Println("a >= b", a >= b)
	fmt.Println("a <= b", a <= b)
	fmt.Println("a != b", a != b)
}

// 逻辑运算
func LogicCase() {
	var a = true
	var b = false
	fmt.Println("a && b", a && b)
	fmt.Println("a || b", a || b)
	fmt.Println("!(a && b)", !(a && b))
}

// 位运算
func BitCase() {
	var a uint8 = 60
	var b uint8 = 13
	var c uint8 = 0
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()

	c = a & b
	fmt.Println("a & b :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()

	c = a | b
	fmt.Println("a | b :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()

	c = a ^ b
	fmt.Println("a ^ b :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()

	c = a << 2
	fmt.Println("a << 2 :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", c)
	fmt.Println()

	c = a >> 2
	fmt.Println("a >> 2 :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", c)
	fmt.Println()

	c = ^a
	fmt.Println("^a :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", c)
	fmt.Println()

	c = a &^ b
	fmt.Println("a &^ b :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()

	c = a & ^b
	fmt.Println("a & ^b :")
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", b)
	fmt.Printf("%08b\n", c)
	fmt.Println()

}

// 赋值运算
func AssignmentCase() {
	var a = 21
	var c int
	c = a
	fmt.Println(" c = a, c值为：", c)
	c += a
	fmt.Println(" c += a, c值为：", c)
	c -= a
	fmt.Println(" c -= a, c值为：", c)
	c *= a
	fmt.Println(" c *= a, c值为：", c)
	c /= a
	fmt.Println(" c /= a, c值为：", c)
	c %= a
	fmt.Println(" c %= a, c值为：", c)

	var b uint8 = 60
	fmt.Printf("b 值为 %d,二进制表示：%08b\n", b, b)
	b <<= 2
	fmt.Printf("b <<= 2, b 值为 %d,二进制表示：%08b\n", b, b)
	b >>= 2
	fmt.Printf("b >>= 2, b 值为 %d,二进制表示：%08b\n", b, b)
	b &= 2
	fmt.Printf("b &= 2, b 值为 %d,二进制表示：%08b\n", b, b)
	b ^= 2
	fmt.Printf("b ^= 2, b 值为 %d,二进制表示：%08b\n", b, b)
	b |= 2
	fmt.Printf("b |= 2, b 值为 %d,二进制表示：%08b\n", b, b)

}
