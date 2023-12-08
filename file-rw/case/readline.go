// @Author huzejun 2023/12/9 4:36:00
package _case

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 以README文件为读取对象
const README = "README.md"

// 一次性读取文件
// 按行拆分打印
// 适合小文件读取
func ReadLine1() {
	fileHandler, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatalln(err)
	}
	defer fileHandler.Close()

	bytes, err := io.ReadAll(fileHandler)
	if err != nil {
		log.Fatal(err)
	}

	list := strings.Split(string(bytes), "\n")
	for _, l := range list {
		fmt.Println(l)
	}
}

// 通过bufio按行读取
// bufio 通过对io模块的封装，提供了对数据的缓冲功能，能以一定程度上减少大数据读写带来的开销
// 当发起读写操作时，会偿试从缓冲区读取数据，缓冲区没有数据后，才会从数据源获取
// 缓冲区大小默认为4k
func ReadLine2() {
	fileHandler, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatalln(err)
	}
	defer fileHandler.Close()
	reader := bufio.NewReader(fileHandler)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}

// 通过 scanner 按行读取
// 单行默认大小 64k
func ReadLine3() {
	fileHandler, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatalln(err)
	}
	defer fileHandler.Close()
	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
