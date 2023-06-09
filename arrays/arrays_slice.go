package main

import "fmt"

func testFunction(arr [5]int) {
	arr[0] = -91
	fmt.Println(arr)

}

func sumFunction(x int) (total int, arr [10]int) {
	for i := 0; i < x; i++ {
		arr[i] = i + 1
		total += arr[i]
	}
	return
}

func main() {
	// mixing of data type in array in Go is not
	var arr [3]int // int array of length 3
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println(arr)

	// creating array using short-hand declaration
	temp := [4]int{1, 2, 3}
	fmt.Println("Array using the shorthand declaration: ", temp)

	// ignoring the length of the Array, this is done by using ... 3 dots like
	arr1 := [...]int{1, 2, 3, 4, 5} // thus here the size of the array becomes dynamic
	fmt.Println("size of arr1: ", len(arr1))

	// arr = arr1 // this will say that can't assgin 5[int] to 3[int] because they are of distinct types..

	// ** Not Reference Type
	// Un-Like Java the Array in Go is Value type not reference type.. meaning if we assign a array to another varaible
	// then a new copy of the original array will be assigned to that variable
	test_array := [...]string{"Microsoft", "Google", "Atlassian", "Visa", "Amazon"}
	gy := test_array
	gy[2] = "Apple"
	fmt.Println("test_array: ", test_array, "\n gy: ", gy)

	//similary array passed to funtion as Parameter not as the reference, the orignal array renmains unchaged
	num := [5]int{1, 2, 3, 4, 5}
	testFunction(num)
	fmt.Println("After function returns the value: ", num)

	// other shortcut way to iterate in Array in Go
	for itr, val := range num { // so here range array-name retruns the index and the value at that index..
		fmt.Printf("%vth index val: %v\n", itr, val)
	}
	val, uparr := sumFunction(10)

	fmt.Println("Sum of 10 natural number: ", val)
	fmt.Println("Array of 10 natural number: ", uparr)

	// if we want only the value not the index then  we can follow the below step..
	// the for loop below will ignore the index
	for _, val := range uparr { // the range upaar will return the value and the index
		fmt.Println(val)
	}

}
