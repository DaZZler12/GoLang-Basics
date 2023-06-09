package main

import (
	"fmt"
	"math"
	"runtime"
)

func forLoopHelper(x int) {
	sum := 0
	for i := 1; i <= x; i++ {
		sum += i
	}
	// fmt.Printf("Sum of %v natural number is %v\n", x, sum)
	fmt.Println("sum of ", x, " and is :", sum)

	// the initialization and increment is Optional
	for sum <= 60 {
		fmt.Println(sum)
		sum++
	}
}

// function with variable length argument along with named return value
func variableArgument(defaultval int, params ...int) (sum int) {
	x := 0
	for x < len(params) { // it's the go's while loop
		sum += params[x]
		x++
	}
	return // implecitly the sum value will be returned
}

func infinteloop() {

	for {
		fmt.Println("lala ")
	}
}

func ifFunction(x, y float64) float64 {
	if temp := math.Sqrt(x); temp < y { // here we can initialize the value before applying the if condition
		return temp
	} else {
		y *= 10
	}
	return y
}

func switchFunction() {
	fmt.Print("Go is running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		break
	case "linux":
		fmt.Println("Linux.")
		break
	default:
		fmt.Printf("%s.\n", os)
	}
}

func stackingdeferFunction() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // this function calls mget stored in the stack and once the loop ends then it will be emptied
	}
	fmt.Println("done")
}

func main() {
	// for loop for sum of n natural number
	forLoopHelper(10)
	fmt.Println("ans from variableArgument method: ", variableArgument(10, 20, 30, 40))
	fmt.Println(ifFunction(2, 16))
	switchFunction()

	// demo of stacking defer in go
	fmt.Println("Demo code defer in Go")
	stackingdeferFunction()
}
