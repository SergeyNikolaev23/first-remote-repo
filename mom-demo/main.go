package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, -15, 0, -16}
	mappedArray := filterNumbers(arr, func(x int) bool {
		return x > 0
	})
	//fmt.Println(mappedArray)
	fmt.Println(arr)
	fmt.Println("Задача 1 ", mappedArray)

	arrTemperatures := []float64{0, 20, 100, -10}
	changeTemp := convertTemperatures(arrTemperatures, func(x float64) float64 {
		return float64((x * 9 / 5) + 32)
	})
	fmt.Println("Задача 2 ", changeTemp)

	arrString := []string{"hello", "world", "go", "programming"}
	arrLength := stringLengths(arrString, func(s string) int {
		return len(s)
	})
	fmt.Println("Задача 3 ", arrLength)
}

type filterFunc func(x int) bool
type convertFunc func(x float64) float64

type lengthFunc func(string) int

func filterNumbers(arr []int, f filterFunc) []int {
	newarr := []int{}
	for i := 0; i < len(arr); i++ {
		if f(arr[i]) {
			newarr = append(newarr, arr[i])
		}
	}
	return newarr
}

func convertTemperatures(arr []float64, f convertFunc) []float64 {
	for i := 0; i < len(arr); i++ {
		arr[i] = f(arr[i])
	}
	return arr
}

func stringLengths(arr []string, f lengthFunc) []int {
	newarr := []int{}
	for i := 0; i < len(arr); i++ {
		newarr = append(newarr, f(arr[i]))
	}
	return newarr
}
