// 1
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var a int = 100
	errors.New(fmt.Sprintf("1"))
	errors.New(fmt.Sprintf("%d", a))
	fmt.Fprintf(os.Stdout, "%d", a)

}
