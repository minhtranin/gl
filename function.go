package main

import "fmt"

func hello() {
	fmt.Println("hello function")
}

func hello1(name string) {
	fmt.Println("hello", name)
}

func hello2(name string) int{
	return 971725697
}

// can return 1 in 2 type
func multi(name, address string) (string int){
	return 4
}

// return 2 argument
func multiReturn(a, b string) (string, int)  {
	return b, 10
}

// named function to straighforward
func named(w, h int) (width int, height int, isEqual bool) {
	isEqual = w == h
	return w, h, isEqual
}

func main() {
	hello()
	hello1("Minh")
	fmt.Println(hello2(""))

	name := multi("M", "TB")
	fmt.Println(name)
	a1, a2 := multiReturn("19", "10")
	fmt.Println(a1, a2)

	_, _, isE := named(10, 10)
	fmt.Println(isE)
}

//
