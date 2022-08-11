package main

import "fmt"

const (
	first = iota + 765
	second
)

func arryDeclartionMethodOne() {

	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 31
	myArray[2] = 8

	fmt.Printf("Arr - %d \n", myArray)
	arryDeclartionMethodTwo()

}

func arryDeclartionMethodTwo() {
	arr := [3]string{"satyan", "Sharma", "jabalpur waale"}
	fmt.Printf("Value of String array is %s \n", arr)

	slice := arr[:]

	fmt.Printf("Value of String slice is %s \n", slice)

	arr[0] = "Ram"

	fmt.Printf("Value of String slice after update is %s \n", slice)

	newSlice := []string{"A", "B", "Hi", "I am", "Learning GOO"}

	fmt.Printf("Value of new slice %s\n", newSlice)

	slice = append(slice, "That's nice ", "Cheers!!!")
	fmt.Printf("after append slice value %s\n", slice)

	s1 := slice[1:]
	s2 := slice[2:3]
	s3 := slice[:3]

	fmt.Println(s1, s2, s3)

}

func understandingMaps() {
	mapOfStrings := map[string]string{"Satyan": "engineer", "Aarna": "Baby"}

	fmt.Printf("Value of  maps-  %s \n", mapOfStrings)
}

func main() {
	fmt.Println("Hello, Satyan")

	/**
	Foloowing Declaration and Usage of a pointer variable
	*/

	var name *string = new(string)
	*name = "Satyan"
	fmt.Println("And the name is - " + *name)

	var age *int = new(int)
	*age = 36
	fmt.Printf("And the age is - %d \n", *age)

	fmt.Printf("Value of constant first %d and second %d \n ", first, second)

	arryDeclartionMethodOne()

	understandingMaps()
	fmt.Println("***End of program***")
}
