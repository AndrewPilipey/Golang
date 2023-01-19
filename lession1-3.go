package main

import "fmt"

func main() { // The clock hand has turned from 12 o'clock by a few degrees. How much time?
	var d, h, m int
	fmt.Scan(&d)
	h = d / 30
	m = d % 30 * 2

	fmt.Println("It is", h, "hours", m, "minutes.")
}
