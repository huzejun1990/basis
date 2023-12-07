// @Author huzejun 2023/12/6 5:39:00
package main

import (
	_case "basis/channel-select/case"
	"os"
	"os/signal"
)

func main() {
	//_case.Communication()
	//_case.ConcurrentSync()
	_case.NoticeAndMultiplexing()
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
