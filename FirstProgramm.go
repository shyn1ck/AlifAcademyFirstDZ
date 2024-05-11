package main

import "fmt"

func main() {
	fmt.Println("Добро пожаловать в мою первую программу. Введите ваше имя:")
	var name string
	fmt.Scan(&name)
	fmt.Println("Здравствуйте,", name)
}
