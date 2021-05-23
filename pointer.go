package main
import "fmt"


func checkPointer(pt *int) {
    fmt.Println(pt)
    fmt.Println("check")
}

func main() {
    fmt.Println("----")
    a := 100
    fmt.Println(a)
    pointer := &a
    fmt.Println(pointer)

    variable := 100
    getPointer := new(int)

    getPointer = &variable
    fmt.Println(*getPointer)

    var getPointer2 *int
    getPointer2 = &variable
    fmt.Println(variable)
    fmt.Println(getPointer2)
    checkPointer(getPointer2)
}