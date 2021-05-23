package main 
import "fmt"

type Animal interface {
    fly()
    speak()
}

type Vehical interface {
    run()
    sound()
}
//embed interface
type SupperAnimal interface {
    Animal
    Vehical
} 

type Dog struct {

}

func(d Dog) speak() { /// in file struct.go.......
    fmt.Println("go go")
}
 
func(d Dog) fly() { /// in file struct.go.......
    fmt.Println("o` o`")
}

func(d Dog) run() { /// in file struct.go.......
    fmt.Println("o` o`")
}

func(d Dog) sound() { /// in file struct.go.......
    fmt.Println("o` o`")
}
 
func goPro(i interface{}) {
    fmt.Println(i)
}

func main () {
   dog := Dog {
   }
   dog.speak()   
   var spAnm SupperAnimal = dog
   spAnm.speak()
   goPro("10")
}