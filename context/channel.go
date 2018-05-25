package main

import (
	"fmt"
	"time"
)

func main() {
	var stop chan bool
	stop = make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Printf("stop~bye~bye~")
				return
			default:
				fmt.Println("just do it~~")
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()

	time.Sleep(time.Second * 3)
	stop <- true
	fmt.Println("stop signal send to go rotine")
	time.Sleep(time.Second * 3)
}
