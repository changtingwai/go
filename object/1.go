// 1
package main

import (
	"fmt"
)

type Integer int

func (a Integer) less(b Integer) bool {
	return a < b
}
func (a *Integer) add(b Integer) {
	*a += b
}
func main() {
	var a Integer = 0
	var flag bool = a.less(1)
	fmt.Println(flag)
	a.add(2)
	fmt.Println(a)
}
