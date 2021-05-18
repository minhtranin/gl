package main
import "fmt"

func main() {
	f1 := make(map[string]int, 2)	
	var f2 map[string]int
	fmt.Println(f1)
	fmt.Println(f2)

	if f1 == nil {
		fmt.Println("# nodejs") // not nil since use make
	}

	if f2 == nil {
		fmt.Println("# nodejs")
	}

	// declare with map

	vMap := map[string]int {
		"Minh": 1,
		"Pro": 2,
	}

	delete(vMap, "Minh")
	vl, check := vMap["Minh1"]
	fmt.Println(vl)
	fmt.Println(check)
	if check {
		fmt.Println("yes")
	}
}