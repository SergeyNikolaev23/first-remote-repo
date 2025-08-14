package main

import "fmt"

// возвращает сумму двух целых чисел a и b
func sum(a, b int) int {
	return a + b
}

func main() {
	result := sum(3, 5)           // вызываем функцию sum() и сохраняем в переменную result
	fmt.Println("Сумма:", result) // выводим на экран
}
