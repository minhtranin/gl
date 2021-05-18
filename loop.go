package main

import "fmt"

func main() {
	variable1 := 10
	for i := 0; i < variable1; i++ {
		fmt.Println(i)
	}
	start := 0
	for ; start < 2; {
		start++
		fmt.Println(start)
	}
	// loop 2
	for i, j := 1, 2; i < 3 && j < 4; i, j = i + 1, j + 1 {
		fmt.Println (i, j)
	}
}

//
