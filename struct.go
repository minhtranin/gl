package main
import (
    "fmt"
    )

type Animal struct {
    name string
    phone int
    call func()
    otherd *string 
}

func (n Animal) getName() string {
        return n.name
    }


func (n *Animal) setName() {
        n.name = "Minh"
    }

func main () {
    otherd := "something else"
    fmt.Println("I ")
    pig := Animal{
        name: "cute",
        phone: 9928,
        call: func() {
            fmt.Println("check function")
        },
        otherd: &otherd,
    }
    fmt.Println(pig)

    var student Animal = struct {
        name string
        phone int
        call func()
        otherd *string 
    }{
        name: "TCM",
        phone: 999,
    }
    pig.call()
    Sname := student.getName()
    fmt.Println(student)
    fmt.Println(Sname)

    // method
    student.setName()
    fmt.Println(student.name)
}