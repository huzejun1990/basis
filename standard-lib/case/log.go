// @Author huzejun 2023/11/19 3:08:00
package _case

import (
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Llongfile)
	log.SetOutput(os.Stderr)
}

func LogCase() {
	var a, b = -1, -2
	_, err := sum(a, b)
	if err != nil {
		log.Println(err)
	}
	log.Printf("a:%d,b%d,两数求和出现错误：%s \n", a, b, err.Error())
	// fatal 日志会导致应用程序的退出
	//log.Fatalf("a:%d,b%d,两数求和出现错误：%s \n",a,b,err.Error())
}
