package main

import "fmt"

func main() {
	food := [4]int {1,2,3,4}
	fmt.Println(food) 
	// without declare size
	box := [...]int {1, 3, 5, 7}
	fmt.Println(box)
	item1 := "MYOB"
	item2 := "else"
	box2 := [...]string {
		item1,
		item2,
	}
	fmt.Println(box2)

	for index, value := range box2 {
		fmt.Println(value)
		fmt.Printf("i=%d value=%s", index, value)
	}

	// array 2 direction
	matrix := [2][4]int {
		{1,2,3,4},
		{5,6,7,8},
	}
	fmt.Println(matrix[0][0])
}

//
