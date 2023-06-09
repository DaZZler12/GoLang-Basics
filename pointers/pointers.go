package main

import "fmt"

func zerovalpointerFun() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}
}

func Dereferencingfun() {
	b := 255
	a := &b
	fmt.Println("address of b is", a)
	fmt.Println("value of b is", *a)
}
func change(val *int) {
	*val = 55
}
func hello() *int {
	i := 5
	return &i
}
func main() {
	fmt.Println("Pointers in Go")
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	// Zero value of a pointer
	zerovalpointerFun()

	// Creating pointers using the new function
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)

	// Dereferencing a pointer
	Dereferencingfun()

	// Passing pointer to a function
	aa := 58
	fmt.Println("value of a before function call is", a)
	bb := &aa
	change(bb)
	fmt.Println("value of a after function call is", a)

	// Returning pointer from a function
	d := hello()
	fmt.Printf("Value of %d, and address: %v\n", *d, d)

	// Go does not support pointer arithmetic
	// b := [...]int{109, 110, 111}
	// p := &b
	// p++  --> this is not possible
}
