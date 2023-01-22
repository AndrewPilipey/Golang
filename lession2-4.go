package main

import "fmt"

// TASK:
//По данному трехзначному числу определите, все ли его цифры различны.
//	Формат входных данных:
//На вход подается одно натуральное трехзначное число.
//	Формат выходных данных:
//Выведите "YES", если все цифры числа различны, в противном случае - "NO".

func main() {
	var a, first, second, third int
	fmt.Println("Enter a three-digit number")
	fmt.Scan(&a)

	first = a / 100
	second = a / 10 % 10
	third = a % 10

	switch {
	case first != second && second != third && third != first:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
}
