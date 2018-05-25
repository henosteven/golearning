package main

import (
	"context"
	"fmt"
	"time"
)

func watch2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("get~stop~signal" +  ctx.Value("flag").(string))
			return
		default:
			fmt.Println("just do it ~" + ctx.Value("flag").(string))
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(ctx.Deadline()) //0001-01-01 00:00:00 +0000 UTC false
	//false时表示没有设置截止时间，如果需要取消的话，需要调用取消函数进行取消。

	ctxv := context.WithValue(ctx, "flag", "info1")
	go watch2(ctxv)

	ctxv = context.WithValue(ctx, "flag", "info2")
	go watch2(ctxv)

	ctxv = context.WithValue(ctx, "flag", "info3")
	go watch2(ctxv)

	ctxv = context.WithValue(ctx, "flag", "info4")
	go watch2(ctxv)

	time.Sleep(time.Second * 3)
	cancel()
	fmt.Println("send the stop signal")
	time.Sleep(time.Second * 3)
}