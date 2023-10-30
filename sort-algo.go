package main

import (
	"math"
)

func main() {

}

// CombSort сортирует используя алгоритм расчесывания
func CombSort[T int | float32 | float64](list []T) []T {
	decreaseFactor := 1.247
	interval := int(math.Floor(float64(len(list)) / decreaseFactor))

	for interval > 0 {
		for i := 0; i+interval < len(list); i++ {
			if list[i] > list[i+interval] {
				list[i], list[i+interval] = list[i+interval], list[i]
			}
		}
		interval = int(math.Floor(float64(interval) / decreaseFactor))
	}

	return list
}

// InsertionSort сортирует используя алгортим вставок
func InsertionSort[T int | float32 | float64](list []T) []T {
	for i := 1; i < len(list); i++ {
		x := list[i]
		j := i

		for j > 0 && list[j-1] > x {
			list[j] = list[j-1]
			j -= 1
		}

		list[j] = x
	}

	return list
}

// SelectionSort сортирует используя алгортим выбора
func SelectionSort[T int | float32 | float64](list []T) []T {
	for i := 0; i < len(list); i++ {
		minElementIndex := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[minElementIndex] {
				minElementIndex = j
			}
		}

		if minElementIndex != i {
			list[i], list[minElementIndex] = list[minElementIndex], list[i]
		}
	}

	return list
}
