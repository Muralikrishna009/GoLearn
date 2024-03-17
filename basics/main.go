package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	fmt.Println("------------")

	fmt.Println("Data Types!")

	var myVar string = "first one"

	fmt.Println(myVar)
	myVar += "new sting added to empty"

	fmt.Println(myVar)

	var firstInt int
	fmt.Println(firstInt)
	fmt.Println("the above was int declaration without any specification")

	var secondInt int8 = 127
	fmt.Println(secondInt)

	fmt.Println("add 1 to int8 number which has the current value of 127")

}
