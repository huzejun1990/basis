// @Author huzejun 2023/12/26 2:41:00
package main

import (
	_case "basis/sync-WaitGroup/case"
	"context"
	"os"
	"os/signal"
)

func main() {
	//_case.WaitGroupCase()
	//_case.WaitGroupCase1()
	//_case.CondCase()
	_case.CondQueueCase()

	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	<-ctx.Done()
}
