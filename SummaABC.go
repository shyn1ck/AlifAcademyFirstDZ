package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	var a, b, c int
	fmt.Println("Консольная программа для нахождение суммы a, b, c")
	fmt.Println("Введите a, b, c")
	fmt.Scan(&a, &b, &c)
	fmt.Println("Сумма a + b + c =", a+b+c)
}
