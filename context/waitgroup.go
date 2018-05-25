package main

import (
	"sync"
	"fmt"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("start~work")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("finish~work")
}