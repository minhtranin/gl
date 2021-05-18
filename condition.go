package main

import "fmt"

func main() {
	money := 2000
	if money < 1000 {
		fmt.Println("impoverish")
	} else {
		fmt.Println("rich man")
	}

	// switch case

	heal := 200
	switch heal {
	case 100, 101, 102, 103, 104:
		fmt.Println("Congratulation!")
		fallthrough
		// without break
	case 200:
		fmt.Println("Good luck")
		fallthrough
	default:
		goto handleGoto
		fmt.Println("Default")
		handleGoto : fmt.Println("this is goto")
	}
}

//
