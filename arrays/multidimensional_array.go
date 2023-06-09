package main

import (
	"fmt"
)

func multidarrayFuntion(multidarr [][]int) (sum int) {
	for index, arr := range multidarr {
		for _, val := range arr {
			sum += val
		}
		fmt.Println((index + 1), "st row:", arr)
	}
	return
}

func main() {
	fmt.Println("MultiDimensional Array")
	multi_arr := [][]int{ // basically this is a multid array whose no of rows is not fixed but the no. of cloumns is always need to be defined
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	fmt.Println("multi: ", multi_arr)
	fmt.Println("multi len/row: ", len(multi_arr))       //return the no. of row
	fmt.Println("multi len/column: ", len(multi_arr[0])) //return the no. of row
	fmt.Println("sum of value of multi_arr", multidarrayFuntion(multi_arr))
}
