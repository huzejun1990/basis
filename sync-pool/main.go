// @Author huzejun 2023/12/22 5:38:00
package main

import (
	_case "basis/sync-pool/case"
	"context"
	"os"
	"os/signal"
)

func main() {
	//_case.PoolCase()
	_case.OnceCase()

	ctx, stop := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt)
	defer stop()
	<-ctx.Done()
}
