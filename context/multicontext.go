package main

import (
	"fmt"
	"time"
	"context"
)

func watch(ctx context.Context, flag string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("get~stop~signal" +  flag)
			return
		default:
			fmt.Println("just do it ~" + flag)
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go watch(ctx, "info1")
	go watch(ctx, "info2")
	go watch(ctx, "info3")
	go watch(ctx, "info4")

	time.Sleep(time.Second * 3)
	cancel()
	fmt.Println("send the stop signal")
	time.Sleep(time.Second * 3)
}