// sorted.go
package main

import (
	"flag"
	"fmt"
)

var infile *string = flag.String("i", "infile", "file contains value for sorting")
var outfile *string = flag.String("o", "outfile", "file to recieve sorted values")
var algorithm *string = flag.String("a", "qsort", "sort algorithm")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile=", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
	}

}
