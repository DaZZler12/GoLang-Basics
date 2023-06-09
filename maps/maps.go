package main

import "fmt"

// data structure to stire data as key value pair
func main() {
	// syntex to make map:
	employeeSalary := make(map[string]int)
	fmt.Println(employeeSalary)

	employeeSalary["steve"] = 12000
	employeeSalary["jamie"] = 15000
	employeeSalary["mike"] = 9000
	fmt.Println("employeeSalary map contents:", employeeSalary)

	// Zero value of a map
	var employeeSalary1 map[string]int
	employeeSalary["steve"] = 12000
	fmt.Println(employeeSalary1)

	// Retrieving value for a key from a map
	employee := "jamie"
	salary := employeeSalary[employee]
	fmt.Println("Salary of", employee, "is", salary)

	// Checking if a key exists
	newEmp := "joe"
	value, ok := employeeSalary[newEmp]
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
		return
	}
	fmt.Println(newEmp, " --> key not found")

	// Iterate over all elements in a map
	fmt.Println("Contents of the map: here we are iterating over the map")
	for key, value := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	}

	// Deleting items from a map
	fmt.Println("map before deletion", employeeSalary)
	delete(employeeSalary, "steve")
	fmt.Println("stev is deleted: map after deletion", employeeSalary)

	// Map of structs
	type employees struct { // defining a struct
		salary  int
		country string
	}

	emp1 := employees{
		salary:  12000,
		country: "USA",
	}
	emp2 := employees{
		salary:  14000,
		country: "Canada",
	}
	emp3 := employees{
		salary:  13000,
		country: "India",
	}
	employeeInfo := map[string]employees{
		"Steve": emp1,
		"Jamie": emp2,
		"Mike":  emp3,
	}

	// iterating over the map of strcuts
	fmt.Println("Iterating over the map of strcuts")
	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.salary, info.country)
	}

	//Length of the map
	fmt.Println("length is", len(employeeInfo))

}
