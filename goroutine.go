package main
import (
    "fmt"
    // "time"
    "runtime"
    "sync"
)
func run1() {
    fmt.Println("run 1")
    wg.Done()
}

func run2() {
    fmt.Println("run 2")
    wg.Done()
}

var wg sync.WaitGroup

func main() {
     go run1()
     ng := runtime.NumGoroutine()
    //  time.Sleep(time.Second)
    fmt.Println(ng)

    // sync go routine
    wg.Add(2)
    fmt.Println("begin...")
    go run1()
    go run2()
    wg.Wait()
    fmt.Println("end...")
}