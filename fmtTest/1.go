// 1
package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string = "jd_a"
	c := strings.Split(a, "_")
	fmt.Print(c[0])
}
