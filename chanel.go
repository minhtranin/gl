package main
import (
    "fmt"
)

func main() {
    ch := make(chan int, 2)
    ch <- 100
    ch <- 1
    close(ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)

    // select in chanel
    queue := make(chan int)
    check := make(chan bool)
    go func(){
        for i :=0; i <10; i++ {
            queue <- i
        }
        check <- true
    }()

    for {
        select {
            case v := <-queue:
                fmt.Println(v)
            case <-check:
                return
        }
    } 
}