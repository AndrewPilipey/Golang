package main

import "fmt"

func main() {
	var last_el int = 123
	last_el = last_el % 10
	fmt.Println(last_el) //display the last element of an integer variable

	var penultimate_el int
	fmt.Println("Enter a number not exceeding 10000")
	fmt.Scan(&penultimate_el)
	penultimate_el = penultimate_el % 100 / 10
	fmt.Println(penultimate_el) //display the penultimate element of an integer variable
}
