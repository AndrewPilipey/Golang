package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Your name is:", name)
	fmt.Println("Your age is:", age)
	fmt.Println("I returned to study Golang")
}
