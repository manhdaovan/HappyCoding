package main

import "fmt"

// struct
// A struct is a collection of fields.
// type StructName struct {
// 	  structField type
//    ...
// }
type Employee struct {
	Name string
	Age uint8
	Department *DepartmentDetail
	Department2 DepartmentDetail
}

type DepartmentDetail struct {
	Name string
	Description string
}

func (e EmpEmployee) IsOver35() bool {
	return e.Age > 35
}

func (e *Employee) ChangeDepartment(newDept *DepartmentDetail) {
	e.Department = newDept
}

type A struct {
	V int
}

func (a *A) changeWithPtr() {
	a.V = 1
}

func (a A) changeWithoutPtr() {
	a.V = 2
}

func someMethod() {
	a,b := A{}, A{}
	a.changeWithPtr()
	fmt.Println(a) // {1}

	b.changeWithoutPtr()
	fmt.Println(b) // {0}
}

func main() {
	// pointer
	// A pointer holds the memory address of a value.
	// var p *T
	var p *int
	i := 1
	p = &i
	// assign value to i via p
	*p = 3 // "dereferencing" or "indirecting".
	fmt.Println(i) // 3
	fmt.Println(p) // 0xc0000ba020
	fmt.Println(*p) // 3
	// re-assign i value
	i = 8
	fmt.Println(i) // 8
	fmt.Println(*p) // 8

	// construct struct
	emp := Employee{
		"Nguyen Van A", 
		33, 
		&DepartmentDetail{
			"Development", 
			"Some description",
		},
	}
	emp2 := Employee {
		Age: 35, // order of field value
		Name: "Nguyen Van B",
		Department: &Department{
			Description: "Some description", // order of field value
			Name: "Sales",
		},
	}
	// access to field of struct
	emp.Name
	emp.Department.Description
	// => struct is similar to class in Java

	// array
	// var a [n]T
	// size n cannot be changed!
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a2 := [3]int{1,2,3}
	var employees [100]*Employee
	matrix := [10][10]int{}
	// loop array
	arr := [10]int{}
	for i, v := range arr {
		fmt.Println(i) // i from 0 to 9
		fmt.Println(v) // v is 0 (10 times)
	}

	// slice
	// var s []T
	// dynamic size
	// The zero value of a slice is nil.
	var s []DepartmentDetail
	fmt.Println(s) // nil
	s2 := []int{1,2,3,4}
	matrix2 := [][]int{}

	// change element of slices that have same underlying array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s) // [3 5 7]
	s[0] = 123
	fmt.Println(s) // [123 5 7]
	fmt.Println(primes) // [2 123 5 7 11 13]

	// Slice length and capacity
	s3 := make([]int, 0, 10)
	// len(s3) = 0
	// cap(s3) = 10
	for i := 0; i < 15; i ++ {
		s3 = append(s3, 1)
		// or:
		// s3[i] = 1 // error on i >= 10
	}
	
	// loop slice
	slice := make([]int, 3, 10)
	for i, v := range slice {
		fmt.Println(i) // i from 0 to 2
		fmt.Println(v) // v is 0 (3 times)
	}

	// map
	// A map maps keys to values.
	// The zero value of a map is nil.
	// var m map[T]T
	var m map[int]string
	m1 := map[int]string{111: "string1", 1: "string2"}
	m2 := map[int]*Employee{
		0: &Employee{},
		1: &Employee{},
	}
	m3 := make(map[string][]int)
	m4 := make(map[string][10]string, 100)
	// loop map
	for k, v := range m1 {
		fmt.Println(k) // 111 and 1 (not guarantee the order)
		fmt.Println(v) // "string1" and "string2" (corresponds to key k)
	}

	// Insert or update an element in map m:
	m[key] = elem
	// Retrieve an element:
	elem = m[key]
	// or: elem, ok = m[key]
	// Delete an element:
	delete(m, key) // no error even key is not existing in m
}