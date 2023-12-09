// @Author huzejun 2023/12/9 18:40:00
package main

import (
	_case "basis/sync/case"
	"context"
	"os"
	"os/signal"
)

func main() {

	//_case.AtomicCase()
	//_case.AtomicCase1()
	_case.AtomicCase2()

	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	<-ctx.Done()
}
