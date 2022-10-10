package main

import "fmt"

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(toString("animal"))
	x, y := multipleReturn(2, 3)
	_, z := multipleReturn(4, 5)
	fmt.Println(x, y, z)
	fmt.Println(variadicFunction(1, 2, 3, 4))
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(variadicFunction(nums...))

	nextInt := initSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	//Anonymous Functions
	func(message string) {
		fmt.Println(message)
	}("Hello World")

}

func add(a int, b int) int {
	return a + b
}

func toString(a string) string {
	return "string - " + a
}

func multipleReturn(a int, b int) (int, int) {
	return a, b
}

// variadic functions
func variadicFunction(nums ...int) int {
	var sum int = 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

// closures
func initSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
