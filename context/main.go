// @Author huzejun 2023/12/10 4:11:00
package main

import (
	_case "basis/context/case"
	"context"
	"os"
	"os/signal"
)

func main() {
	_case.ContextCase()

	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	<-ctx.Done()
}
