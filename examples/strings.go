package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	fmt.Println(len(s))                    //total byte length of a string
	fmt.Println(utf8.RuneCountInString(s)) //number of runes in string
	for i, c := range s {
		fmt.Printf("%#U - %d\n", c, i) //prints the index of rune and unicde representation of rune
	}
	for i, w := 0, 0; i < len(s); i += w {
		char, width := utf8.DecodeRuneInString(s[i:]) //decode the rune
		fmt.Printf("%#U\n", char)
		if char == 'ส' {
			fmt.Printf("Found character\n")
		}
		w = width
	}
}
