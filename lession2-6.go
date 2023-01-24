package main

import "fmt"

func main() {
	var win int
	fmt.Println("Enter a number of your ticket - ******")
	fmt.Scan(&win)

	a1 := win / 100000
	a2 := win / 10000 % 10
	a3 := win / 1000 % 10
	sum_a := a1 + a2 + a3

	b1 := win / 100 % 10
	b2 := win / 10 % 10
	b3 := win % 10
	sum_b := b1 + b2 + b3

	if sum_a == sum_b {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/*Формат входных данных
На вход подается номер билета - одно шестизначное  число.

Формат выходных данных
Выведите "YES", если билет счастливый, в противном случае - "NO".

Sample Input:
613244
Sample Output:
YES*/
