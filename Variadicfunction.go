package main
import "fmt"
func variadic(quantity int, food ...string) ([]string) {
	return food
}
func main() {
	arrayFood := variadic(3, "T", "C")
	fmt.Println(arrayFood)

	slice := []string {"P", "R", "O"}
	arrayFood2 := variadic(4, slice...)
	fmt.Println(arrayFood2)
}