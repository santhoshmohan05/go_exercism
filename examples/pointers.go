package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroValPtr(ival *int) {
	*ival = 0
}

func main() {
	var i int = 5
	zeroVal(i)
	fmt.Println(i)
	zeroValPtr(&i)
	fmt.Println(i)
}
