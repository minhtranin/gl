package main

import "fmt"
import "math"

func main() {
	var number int
	number = 10
	fmt.Println(number)
	var email string
	email = "minhtran.in@outlook.com"
	fmt.Println(email)
	// declare 2 variables
	var a, b int
	a = 1
	b = 1
	fmt.Println(a)
	fmt.Println(b)
	// declare and gant

	var c, d int = 2, 2
	fmt.Println(c)
	fmt.Println(d)

	// declare variable without type

	var e, f = 3, 3
	fmt.Println(e)
	fmt.Println(f)

	// declare difference of type

	var (
		name string
		address string
		phone int = 971725797
	)
	name = "Minh T"
	address = "Q TAN BINH"
	fmt.Println(name + address )
	fmt.Println(phone )
	
	// declare 1 line
	var first, second, third = "one", "two", 3  
	fmt.Println(first + second )
	fmt.Println(third )


	// without var
	finalize := "done"
	fmt.Println(finalize)

	// type boolean

	var myBool bool = true
	newBool := false
	fmt.Println(myBool)
	fmt.Println(newBool)

	// range max int
	fmt.Println(math.MinInt8)
	fmt.Println(math.MinInt16)
	fmt.Println(math.MinInt32)
	fmt.Println(math.MinInt64)
	complexV1 := 10 + 1i
	complexV2 := 10 + 1i
	fmt.Println(complexV1 + complexV2)
	
	// type rune
	spell := "TCM"
	runes := []rune(spell)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}

	testRun := "Trần Công Minh"
	fmt.Println(testRun)
}

//
