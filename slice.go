package main
import "fmt"

func main()  {
	var variable1 []int
	fmt.Println(variable1)

	const variable2 string = "Minh"
	fmt.Println(variable2)

	variable3 := "Test"
	variable3 = "changing"
	fmt.Println(variable3) 

	// slice
	variable4 := [] int {1,2,3,4}
	fmt.Println(variable4)

	// slice from an array
	minorVariable4 := variable4[0:2] 
	fmt.Println(minorVariable4)

	// cp array
	cpVariable4 := variable4[:]
	fmt.Println(cpVariable4)


	// make copy append
	/// make - provide a paticular cap
	sliceMake := make([]int, 2, 5)
	fmt.Println(sliceMake)
	fmt.Println(len(sliceMake)) // 2
	fmt.Println(cap(sliceMake)) // 5

	// copy slice to cap

	needSlice := []string{"T", "C", "M"}
	inspireMake := make([]string, 2)
	copy(inspireMake, needSlice)
	fmt.Println(inspireMake)

	// append
	newSlice := append(needSlice[:1], needSlice[2:]...)
	fmt.Println(newSlice)
	
}