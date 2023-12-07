// @Author huzejun 2023/12/6 5:40:00
package _case

import (
	"fmt"
	"time"
)

// 协程进通信
func Communication() {
	//定义一个可读，可写的通道
	ch := make(chan int, 0)
	go communicationF1(ch)
	go communicationF2(ch)
}

// F1 接收一个只写通道
func communicationF1(ch chan<- int) {
	// 通过循环向通道写入 0~99
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

// F2 接收一个只读通道
func communicationF2(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

// 并发场景下, 同步机制
func ConcurrentSync() {
	//带缓冲的通道
	ch := make(chan int, 10)
	//向ch 写入消息
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	//向 ch 写入消息
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	// 从ch 读取消息，并打印
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
}

// 通知协程能出与多路复用
func NoticeAndMultiplexing() {
	ch := make(chan int, 0)
	strCh := make(chan string, 0)
	done := make(chan struct{}, 0)
	go noticeAndMultiplexingF1(ch)
	go noticeAndMultiplexingF2(strCh)
	go noticeAndMultiplexingF3(ch, strCh, done)
	time.Sleep(5 * time.Second)
	close(done)

}

func noticeAndMultiplexingF1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}

func noticeAndMultiplexingF2(ch chan<- string) {
	for i := 0; i < 100; i++ {
		ch <- fmt.Sprintf("数字：%d", i)
	}
}

// select 子句作为一个整体阻塞，其中任意channel准备就绪则继续执行
func noticeAndMultiplexingF3(ch <-chan int, strCh <-chan string, done <-chan struct{}) {
	i := 0
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case str := <-strCh:
			fmt.Println(str)
		case <-done:
			fmt.Println("收到退出通知，能出当前协程")
			return
			//default:
			//	fmt.Println("执行default语句")
		}
		i++
		fmt.Println("累计执行次数：", i)
	}
}
