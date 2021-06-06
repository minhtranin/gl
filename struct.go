package main

import (
	"fmt"
)

type Action interface {
	fly()
}

type Animal struct {
	name   string
	phone  int
	call   func()
	otherd *string
}

func (n Animal) getName() string {
	return n.name
}

func (n Animal) fly() {
	fmt.Println(n.name)
}

func (n *Animal) returnInter() Action {
	return Animal{
		name:  "fly return interface",
		phone: 123123,
	}
}
func (n *Animal) setName() {
	n.name = "Minh"
}

func (n *Animal) mm() Animal {
	return Animal{
		name: "pro",
	}
}

func main() {
	otherd := "something else"
	fmt.Println("I ")
	pig := Animal{
		name:  "cute",
		phone: 9928,
		call: func() {
			fmt.Println("check function")
		},
		otherd: &otherd,
	}
	fmt.Println(pig)

	var student Animal = struct {
		name   string
		phone  int
		call   func()
		otherd *string
	}{
		name:  "TCM",
		phone: 999,
	}
	pig.call()
	Sname := student.getName()
	fmt.Println(student)
	fmt.Println(Sname)

	// method
	student.setName()
	fmt.Println(student.name)
	mmstr := student.mm()
	fmt.Println(mmstr)

	var testFly Action = Animal{
		name: "Minh testFly",
	}
	testFly.fly()
	var returnF Action = student.returnInter()
	fmt.Println(returnF.phone)
	fmt.Println(returnF)
}
