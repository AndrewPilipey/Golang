package main

import "fmt"

func main() {
	var year int
	fmt.Println("Enter your year, when are you living in. If you are from 10001 year and more - use another program.")

	fmt.Scan(&year)

	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		fmt.Println("This year is a leap year") // "YES" for the task
	} else {
		fmt.Println("This isn't a leap year") //"NO" for the task
	}
}

/* Требуется определить, является ли данный год високосным, напомним:
Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:
- кратен 400;
- кратен 4, но не кратен 100.

Входные данные
Вводится единственное число - номер года (целое, положительное, не превышает 10000).

Выходные данные
Требуется вывести слово YES, если год является високосным и NO - в противном случае.

Sample Input:
2000

Sample Output:
YES  */
