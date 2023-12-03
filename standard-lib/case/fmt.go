// @Author huzejun 2023/11/14 3:11:00
package _case

import (
	"fmt"
	"os"
)

func FmtCase() {
	// 打印到标准输出
	fmt.Println("今天天气很好")
	//格式化，并打印到标准输出
	fmt.Printf("%s天气很好\n", "今天")
	//格式化
	str := fmt.Sprintf("%s天气很好", "今天,明天")
	//输出到 io.writer
	fmt.Fprint(os.Stderr, str)
	fmt.Println()
}

func FmtCase1() {
	type simple struct {
		value int
	}
	a := simple{
		value: 10,
	}

	// 通用占位符
	fmt.Printf("默认格式的值：%v \n", a)
	fmt.Printf("包含字段名称的值：%+v \n", a)
	fmt.Printf("go语法表示的值：%#v \n", a)
	fmt.Printf("go语法表示的类型：%T \n", a)
	fmt.Printf("输出字面上的百分号：%%10 \n")

	// 整数点位符
	v1 := 10
	v2 := 20170 //"今" 字码点值
	fmt.Printf("二进制：%b \n", v1)
	fmt.Printf("Unicode 码点转字符：%c \n", v2)
	fmt.Printf("十进制：%d \n", v1)
	fmt.Printf("八进制：%o \n", v1)
	fmt.Printf("0o为前缀的八进制：%O \n", v1)
	fmt.Printf("用单引号瘵字符的值包起来：%q \n", v2)
	fmt.Printf("十六进制：%x \n", v1)
	fmt.Printf("十六进制的大写：%X \n", v1)
	fmt.Printf("Unicode格式：%U \n", v2)

	// 宽度设置
	fmt.Printf("%v 的二进制：%b; go语法表示二进制为：%#b; 指定二进制宽度为8，不足8位补8：%08b \n", v1, v1, v1, v1)
	fmt.Printf("%v 的十六进制：%x; 使用go语法表示为,指定十六进制宽度为8，不足8位补8：%#08x \n", v1, v1, v1)
	fmt.Printf("%v 的字符为：%c; 指定宽度为5位，不足5位为空格：%5c \n", v2, v2, v2)

	// 浮点数占位符
	var f1 = 123.789
	var f2 = 12345678910.78999
	fmt.Printf("指数为二的幂的无小数科学计数：%b \n", f1)
	fmt.Printf("使用科学计数法：%e \n", f1)
	fmt.Printf("使用科学计数法,大写：%E \n", f1)
	fmt.Printf("有小数点而无指数，即常规的浮点数据格式，默认宽度和精度：%f \n", f1)
	fmt.Printf("宽度为9，精度默认：%9f \n", f1)
	fmt.Printf("默认宽度为9，精度保留两位小数：%.2f \n", f1)
	fmt.Printf("宽度为9，精度保留两位小数：%09.2f \n", f1)
	fmt.Printf("宽度为9，精度保留0小数：%9.f \n", f1)
	fmt.Printf("根据情况自动选择%%e 或 %%f 来输出，以产生更紧凑的输出（末位无0）：%g %g \n", f1, f2)
	fmt.Printf("根据情况自动选择%%E 或 %%f 来输出，以产生更紧凑的输出（末位无0）：%G %G \n", f1, f2)
	fmt.Printf("以十六进制方式表示：%x \n", f1)
	fmt.Printf("以十六进制方式表示,大写：%X \n", f1)

	// 字符串占位符
	var str = "今天是个好日子"
	fmt.Printf("描述一个今天：%s \n", str)
	//用双引号将字符串包裹
	fmt.Printf("描述一个今天：%q \n", str)
	// 16进制表示
	fmt.Printf("%x \n", str)
	// 以空格作为两数之间的分隔符，并用大写16进制
	fmt.Printf("% X \n", str)

	// 指针点位符
	var str1 = "今天是个好日子"
	bytes := []byte(str1)
	//切片第0个元素的地址
	fmt.Printf("%p \n", bytes)
	mp := make(map[string]int, 0)
	fmt.Printf("%p \n", mp)
	var p *map[string]int = new(map[string]int)

	fmt.Printf("%p \n", p)
}
