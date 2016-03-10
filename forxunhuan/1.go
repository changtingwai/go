// 1
package main

import (
	"errors"
	"fmt"
)

func main() {
	sum := 1
	for {
		if sum > 5 {
			break
		}
		fmt.Println(sum)
		sum++
	}
	fmt.Println("---------------------")

	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
	fmt.Println("---------------------")
''
	b := a[:3]
	fmt.Println(b)
	fmt.Println("---------------------")

	fmt.Println(Add(-5, 6))

}
func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("不支持负数")
		return
	}
	return a + b, nil
}
