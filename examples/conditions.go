package main

import (
	"fmt"
	"time"
)

func main() {
	i := 5
	if i < 0 {
		fmt.Println("Negative ")
	} else if i > 0 {
		fmt.Println("positive")
	} else {
		fmt.Println("zero")
	}
	if num := 9; num < 10 {
		fmt.Println("Less then 10")
	}

	//switch
	j := 15
	switch j {
	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	default:
		fmt.Println("default", j)
	}
	// can call function for switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	case time.Monday:
		fmt.Println("First day of week")
	default:
		fmt.Println("Weekday")
	}
	// switch without case is like if/else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	//switch to compare types
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
