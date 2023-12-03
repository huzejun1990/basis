// @Author huzejun 2023/11/7 3:11:00
package _case

import "fmt"

func FlowControlCase() {
	ifElseCase(0)
	ifElseCase(1)
	ifElseCase(2)
	forCase()
	switchCash("A", 1)
	switchCash("C", "")
	gotoCase()
}

func ifElseCase(a int) {
	if a == 0 {
		fmt.Println("执行 if 语句块")
	} else if a == 1 {
		fmt.Println("执行 else if 语句块")
	} else {
		fmt.Println("执行 else 语句块")
	}
}

func forCase() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println("位置1 执行 for 语句块 i: ", i)
	}

	var i int
	var b = true
	for b {
		fmt.Println("位置2 执行 for 语句块 i: ", i)
		i++
		if i >= 10 {
			b = false
		}
	}

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for index, data := range list {
		fmt.Printf("位置3 执行 for 语句块 index: %d, data: %d \n", index, data)
	}

	mp := map[string]string{"A": "a", "B": "b", "C": "c"}
	for key, value := range mp {
		fmt.Printf("位置4 执行 for 语句块 key: %v, value: %v \n", key, value)
	}

	str := "今天天气很好"
	for index, char := range str {
		fmt.Printf("位置5 执行 for 语句块 index: %v,char : %v \n", index, string(char))
	}

	var j int
	for {
		fmt.Println("位置6 执行 for 语句块 j: ", j)
		j++
		if j >= 10 {
			break
		}
	}

lab1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("位置7 执行 for 语句块 i: %d, j : %d \n", i, j)
			break lab1
		}
	}
}

func switchCash(alpha string, in interface{}) {
	switch alpha {
	case "A":
		fmt.Println("执行 A 语句块")
	case "B":
		fmt.Println("执行 B 语句块")
	case "C", "D":
		fmt.Println("执行 CD 语句块")
		fallthrough
	case "E":
		fmt.Println("执行 E 语句块")
	case "F":
		fmt.Println("执行 F 语句块")
	}

	switch in.(type) {
	case string:
		fmt.Println("in 输入为字符串")
	case int:
		fmt.Println("in 输入为int类型")
	default:
		fmt.Println("未能识别in的类型")

	}

}

func gotoCase() {
	var a = 0
lab1:
	fmt.Println("goto 位置1")
	for i := 0; i < 10; i++ {
		a += i
		if a == 0 {
			a += 1
			goto lab1
		}
	}
}
