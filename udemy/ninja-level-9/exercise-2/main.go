// This exercise will reinforce our understanding of method sets:
// ● create a type person struct
// ● attach a method speak to type person using a pointer receiver
// ○ *person
// ● create a type human interface
// ○ to implicitly implement the interface, a human must have the speak method
// ● create func “saySomething”
// ○ have it take in a human as a parameter
// ○ have it call the speak method
// ● show the following in your code
// ○ you CAN pass a value of type *person into saySomething
// ○ you CANNOT pass a value of type person into saySomething
// ● here is a hint if you need some help
// ○ https://play.golang.org/p/FAwcQbNtMG

package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hello! My Name is", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Hoo",
	}
	saySomething(&p1)
	// cannot use p1 (variable of type person) as type human in argument to saySomething:
	// person does not implement human (speak method has pointer receiver)
	// saySomething(p1)

	// method sets와 interface 구현의 관계를 이해하자!
}
