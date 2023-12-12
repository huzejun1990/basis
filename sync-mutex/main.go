// @Author huzejun 2023/12/12 20:58:00
package main

import (
	_case "basis/sync-mutex/case"
	"context"
	"os"
	"os/signal"
)

func main() {

	//_case.MutexCase()
	//_case.MapCase()
	_case.MapCase1()

	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	<-ctx.Done()
}
