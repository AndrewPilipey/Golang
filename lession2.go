package main

import "fmt"

// convert minutes to hours and minutes
func main() {
	var (
		your_time int
		hours     int
		minutes   int
	)
	fmt.Println("How many minutes?")
	fmt.Scan(&your_time)
	hours = your_time / 60
	minutes = your_time % 60

	fmt.Println("It is", hours, "hours", minutes, "minutes")
}
