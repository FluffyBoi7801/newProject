package main

import (
	"fmt"
)

func main() {
	var a int32

	fmt.Println("Привет! Введи целое число!")
	fmt.Scanf("%d", &a)
	fmt.Printf("Ух ты! Ты ввел: %d", a)
}
