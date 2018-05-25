package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("get~stop~signal")
				return
			default:
				fmt.Println("just do it ~")
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()

	time.Sleep(time.Second * 3)
	cancel()
	fmt.Println("send the stop signal")
	time.Sleep(time.Second * 3)
}