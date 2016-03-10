// 1
package main

import (
	"fmt"
)

func count(ch chan int, i int) {
	ch <- i
	fmt.Println("counting")
}
func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go count(chs[i], i)
	}
	for _, ch := range chs {
		fmt.Println(<-ch)
	}
}
