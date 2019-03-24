package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//var s, sep string

	//another type of for loop that iterates over range
	//for _, arg := range os.Args[0:] {

	//for large data this below could be expensive. hence use strings.Join from strings package
	//s += sep + arg
	//sep = " "
	//}
	// one type of for loop below
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	fmt.Println(strings.Join(os.Args[1:], " "))
}
