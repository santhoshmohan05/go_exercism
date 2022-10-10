package main

import "fmt"

func arrays() {
	//Arrays - Fixed Length
	var a [5]int
	fmt.Println(a)
	a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	var b [3]int = [3]int{7, 8, 9}
	fmt.Println(b)
	// auto initialise
	c := [2]int{0, 1}
	fmt.Println(c)
	var d [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			d[i][j] = i + j
		}
	}
	fmt.Println(d)
}

func slices() {
	// slices - Variable length
	var e []int
	e = []int{1, 2, 3}
	e = append(e, 4)
	fmt.Println(e)
	// using make
	f := make([]string, 3)
	f[0] = "a"
	f[1] = "b"
	f[2] = "c"
	f = append(f, "d")
	fmt.Println(f)
	//using copy
	g := make([]string, len(f))
	copy(g, f)
	fmt.Println(g)
	h := f[1:3]
	i := f[2:]
	j := f[:2]
	fmt.Println(h, i, j)
	// 2D slices
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func maps() {
	a := make(map[string]int)
	a["a"] = 1
	a["b"] = 2
	a["c"] = 3
	fmt.Println(a)
	delete(a, "a")
	fmt.Println(a)
	// map access returns second value to indicate if key is present in map
	value, prs := a["k2"]
	fmt.Println("prs:", prs)
	fmt.Println("prs:", value)
	var n map[int]int = map[int]int{1: 2, 3: 4}
	fmt.Println(n)

}

func main() {
	arrays()
	slices()
	maps()

}
