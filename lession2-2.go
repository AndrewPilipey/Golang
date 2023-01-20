package main

import "fmt"

func main() {
	var a bool = true
	var b bool = !a
	var c bool = !b

	fmt.Println(a, b, c)

	var d bool = 4 > 3 && 10.0 > 5.05
	var e bool = 99 < 1 && 34 > 2

	fmt.Println(d, e)

	var f bool = 10 > 5 || 2 > 1
	var g bool = 45 > 44 || 44 <= 43

	fmt.Println(f, g)

	A := 7
	B := 7
	if A < B {
		fmt.Println("A is less than B")
	} else if A > B {
		fmt.Println("A is more than B")
	} else {
		fmt.Println("A and B are equal")
	}
}
