package main

import (
	"fmt"
	"math"
)

func main() {
	//variables
	var a = "Initial"
	var b = 5.0
	var c int = 1
	var d, e int = 2, 3
	f := true
	fmt.Println(a, b, c, d, e, f)
	//constants
	const g = 50
	const h = 500 / g
	fmt.Println(g, h, math.Sin(h))
}
