package main

import (
	"context"
	"fmt"
	"time"
)

func longProccess(ctx context.Context, ch chan string) {
	fmt.Println("run ")
	time.Sleep(5 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cansel := context.WithTimeout(ctx, 3*time.Second)
	defer cansel()
	go longProccess(ctx, ch) //とんでもなく時間がかかる時はキャンセルしたい

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case <-ch:
			fmt.Println("success")
			return
		}
	}
}
