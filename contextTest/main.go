package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	ctx := context.Background()
	subCtx, cancel := context.WithCancel(ctx)

	go func(ctx context.Context) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("subCtx Done")
			return
		}
	}(subCtx)

	time.AfterFunc(time.Second*2, func() {
		fmt.Println("start cancel")
		cancel()
	})

	wg.Wait()
	fmt.Println("Modified by zsx@20220724")
	fmt.Println("Modified by zsx@20220824")
}
