package main

import "fmt"

func Abc() {

}

type ExampleStruct struct{}

func (a *ExampleStruct) ExportedMethod(i int) (int, error) {
	return i, nil
}

func (a ExampleStruct) unexportedMethod(i int) {

}

type AliasOfInt int

func (b AliasOfInt) ExportedMethod() {

}

// Interface values: (value, type)
// Detail: https://blog.golang.org/laws-of-reflection
type Animal interface {
	Move()
}

type Dog struct{}

func (d *Dog) Move() {
	if d == nil {
		// do something here
	}
}

type Shark struct{}

func (d Shark) Move() {}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	es := ExampleStruct{}
	es.ExportedMethod(1)
	es.unexportedMethod(2)

	var a Animal
	if a == nil {
		// true
	}
	describe(a) // (<nil>, <nil>)

	a = &Dog{}
	describe(a) // (&{}, *main.Dog)
	a.Move()

	a = Shark{}
	describe(a) // ({}, main.Shark)
	a.Move()

	TestInterface() // a is not nil

	// interface{} says nothing
	var i interface{}
	describe(i) // (<nil>, <nil>)

	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)

	TypeAssert(a)

	TypeSwitch(a)

	// error
	type error interface {
		Error() string
	}
}

func GetAnimal() Animal {
	var d *Dog = nil
	// var a Animal
	return d
}

func TestInterface() {
	a := GetAnimal()
	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println("a is not nil")
	}
}

func TypeAssert(a interface{}) {
	d := a.(*Dog)
	d, ok := a.(*Dog)
	if !ok {
		// 
	}
	d.Move()
	fmt.Println(d)
}

func TypeSwitch(a Animal) {
	switch a.(type) {
	// switch v := a.(type) {
	case *Dog:
		fmt.Println("it is a dog")
	case Shark:
		fmt.Println("it is a shark")
	default:
		fmt.Println("undefined animal")
	}
}
