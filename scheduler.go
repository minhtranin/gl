package main 
import (
    "fmt"
    "runtime"
)
func main() {
    // runtime.GOMAXPROCS()
    numC := runtime.NumCPU()
    fmt.Println(numC)

    numM := runtime.GOMAXPROCS(0)
    fmt.Println(numM)
}