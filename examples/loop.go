package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	for j := 2; j <= 5; j++ {
		fmt.Println(j)
	}
	i = 0
	for {
		fmt.Println("endless Loop")
		if i == 0 {
			i++
			continue
		}
		break
	}

	a := []string{"Foo", "Bar"}
	for i, s := range a {
		fmt.Println(i, s)
	}
	b := "wholetext"
	for i, s := range b {
		fmt.Printf("%c is in position %d\n", s, i)
	}
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for key, val := range m {
		fmt.Println(key, val)
	}
}
