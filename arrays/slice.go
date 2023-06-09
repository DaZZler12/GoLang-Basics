package main

import "fmt"

func sliceFunction(nos []int) {
	for index := range nos {
		nos[index] = nos[index] * 2
	}
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func main() {
	var a [5]int // array of length 5
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}
	fmt.Println(a)

	// slice is a wrapper on top of array, slice don't have on their own.. they are just like refernce to am array
	// syntex
	var b []int = a[1:4]                             // here b is a slice on array a having reference form index 1 to 3 of the array a
	fmt.Println("slice of a: start=1 and end=4 ", b) //here what ever end we give it will take the value of end-1

	// a[start:end] ==> index covered is start , end-1

	// slice import
	// so when we write like below so basically we are creating an array and putting the slice to a i.e. we are creating a slice
	// a:=[8]int{

	// }

	// similary the below is the array
	// a:=[][]int{
	// 	{},
	// 	{},
	// 	{},
	// }

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++ // it's modifying the main array as it has the reference
	}
	fmt.Println("array after", darr)

	allslcie := darr[:] //creats a slice with all elements
	fmt.Println(allslcie)

	fmt.Println("len of slice: ", len(dslice))
	fmt.Println("Capacity of slice: ", cap(dslice))

	//create slice using make
	myslice := make([]int, 5, 5) // it creates an array and give the refrence to the slice
	//syntex of make function: make([]T, len, cap) ([]dataType, len , capacity)
	fmt.Println("myslice", myslice)
	// the capacity field is optional

	// Appending to a slice
	cars := []string{"Ferrari", "Honda", "Ford"}                                       // this is a slice
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	cars = append(cars, "Jaguar")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	cars = append(cars, "Benz")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	cars = append(cars, "Ola")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 12

	var names []string //zero value of a slice is nil
	fmt.Println(len(names))
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	//appending 2 slice
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	qunt := []string{"1", "2", "3", "4", "5"}
	food = append(food, qunt...)
	fmt.Println("food:", food)

	// passing slice as the funciton parameter
	nos := []int{10, 8, 6}
	fmt.Println("Before paased to function: ", nos)
	sliceFunction(nos)
	fmt.Println("After function is called: ", nos)

	// multi-dimensional slice
	fmt.Println("MultiD slice: ")
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	countriesNeeded := countries()
	fmt.Println("countriesNeeded: ", countriesNeeded)
}
