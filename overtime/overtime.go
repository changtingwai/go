// overtime
package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	ch := make(chan int, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	ch <- 1
	select {
	case <-ch:
		fmt.Println("Hello World!")
	case <-timeout:
		fmt.Println("timeout!")
	}

}
