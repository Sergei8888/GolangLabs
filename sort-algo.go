package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

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

// ShellSort сортирует используя алгоритм Шелла
func ShellSort[T int | float32 | float64](list []T) []T {
	for gap := len(list) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(list); i++ {
			for j := i; j >= gap && list[j-gap] > list[j]; j -= gap {
				list[j], list[j-gap] = list[j-gap], list[j]
			}
		}
	}

	return list
}

// Finds the largest number in an array
func findLargestNum[T int | float32 | float64](array []T) T {
	var largestNum T = 0

	for i := 0; i < len(array); i++ {
		if array[i] > largestNum {
			largestNum = array[i]
		}
	}
	return largestNum
}

// RadixSort сортирует используя алгоритм поразрядной сортировки
func RadixSort(list []int) []int {
	// Base 10 is used
	largestNum := findLargestNum(list)
	size := len(list)
	var significantDigit = 1
	semiSorted := make([]int, size)

	// Loop until we reach the largest significant digit
	for largestNum/significantDigit > 0 {
		bucket := [10]int{0}

		// Counts the number of "keys" or digits that will go into each bucket
		for i := 0; i < size; i++ {
			bucket[(list[i]/significantDigit)%10]++
		}

		// Add the count of the previous buckets
		// Acquires the indexes after the end of each bucket location in the array
		// Works similar to the count sort algorithm
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}

		// Use the bucket to fill a "semiSorted" array
		for i := size - 1; i >= 0; i-- {
			bucket[(list[i]/significantDigit)%10]--
			semiSorted[bucket[(list[i]/significantDigit)%10]] = list[i]
		}

		// Replace the current array with the semisorted array
		for i := 0; i < size; i++ {
			list[i] = semiSorted[i]
		}

		// Move to next significant digit
		significantDigit *= 10
	}

	return list
}

// HeapSort сортирует методом кучи
func HeapSort(A []int) []int {
	// Build the heap, get it's size
	buildMaxHeap(A)
	size := len(A)

	for size > 1 {
		// Repeatedly swap first and last element, decrement size of considered list
		A[0], A[size-1] = A[size-1], A[0]
		size -= 1

		// SiftDown on first element to put it in place
		siftDown(A[:size], 0)
	}

	return A
}

// Build a Max heap from a given array.
// Max heap adheres to the properties of:
// * Being filled left to right
// * No child node is greater than its parents
func buildMaxHeap(A []int) {
	// Since all leaf nodes are considered valid heaps, we start from the last
	// Non-leaf node, and work right to left in the array
	for i := (len(A) / 2) - 1; i >= 0; i -= 1 {
		siftDown(A, i)
	}
}

// Given an index i of A, sift it downwards, should it be smaller than its children
func siftDown(A []int, i int) {
	largest := i
	l, r := 2*i+1, 2*i+2

	if l < len(A) && A[l] > A[i] {
		largest = l
	}

	if r < len(A) && A[r] > A[largest] {
		largest = r
	}

	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		siftDown(A, largest)
	}
}

// MergeSort сортирует используя алгоритм слияния
func MergeSort[T int | float32 | float64](list []T) []T {
	if len(list) <= 1 {
		return list
	}

	middle := len(list) / 2
	left := MergeSort(list[:middle])
	right := MergeSort(list[middle:])

	return merge(left, right)
}

func merge[T int | float32 | float64](left []T, right []T) []T {
	result := make([]T, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}

		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}

// QuickSort сортирует используя алгоритм быстрой сортировки
func QuickSort[T int | float32 | float64](list []T) []T {
	if len(list) < 2 {
		return list
	}

	left, right := 0, len(list)-1

	pivot := rand.Int() % len(list)

	list[pivot], list[right] = list[right], list[pivot]

	for i := range list {
		if list[i] < list[right] {
			list[left], list[i] = list[i], list[left]
			left++
		}
	}

	list[left], list[right] = list[right], list[left]

	QuickSort(list[:left])
	QuickSort(list[left+1:])

	return list
}

// PolyPhaseMergeSort сортирует используя алгоритм многофазной сортировки
func PolyPhaseMergeSort(inputFiles []string, outputFile string) error {
	const mergeSize = 2 // Number of chunks to merge at a time
	var mergedFiles []string

	for len(inputFiles) > 1 {
		for i := 0; i < len(inputFiles); i += mergeSize {
			chunksToMerge := inputFiles[i:min(i+mergeSize, len(inputFiles))]
			var mergedChunk []int
			file1, err := os.Open(chunksToMerge[0])
			if err != nil {
				return err
			}
			file2, err := os.Open(chunksToMerge[1])
			if err != nil {
				return err
			}

			scanner1 := bufio.NewScanner(file1)
			scanner2 := bufio.NewScanner(file2)

			var line1, line2 string
			if scanner1.Scan() {
				line1 = scanner1.Text()
			}
			if scanner2.Scan() {
				line2 = scanner2.Text()
			}

			for line1 != "" && line2 != "" {
				num1, _ := strconv.Atoi(line1)
				num2, _ := strconv.Atoi(line2)
				if num1 < num2 {
					mergedChunk = append(mergedChunk, num1)
					if scanner1.Scan() {
						line1 = scanner1.Text()
					} else {
						line1 = ""
					}
				} else {
					mergedChunk = append(mergedChunk, num2)
					if scanner2.Scan() {
						line2 = scanner2.Text()
					} else {
						line2 = ""
					}
				}
			}

			for line1 != "" {
				num, _ := strconv.Atoi(line1)
				mergedChunk = append(mergedChunk, num)
				if scanner1.Scan() {
					line1 = scanner1.Text()
				} else {
					line1 = ""
				}
			}

			for line2 != "" {
				num, _ := strconv.Atoi(line2)
				mergedChunk = append(mergedChunk, num)
				if scanner2.Scan() {
					line2 = scanner2.Text()
				} else {
					line2 = ""
				}
			}

			file1.Close()
			file2.Close()

			mergedFile, err := sortAndWriteChunk(mergedChunk)
			if err != nil {
				return err
			}
			mergedFiles = append(mergedFiles, mergedFile)
		}

		inputFiles = mergedFiles
		mergedFiles = nil
	}

	err := os.Rename(inputFiles[0], outputFile)
	if err != nil {
		return err
	}

	return nil
}

// Function to perform external sorting
func externalSort(inputFile string, outputFile string) error {
	const chunkSize = 1000 // Size of each chunk
	chunks := make([][]int, 0)

	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var chunk []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		chunk = append(chunk, num)
		if len(chunk) == chunkSize {
			chunks = append(chunks, chunk)
			chunk = nil
		}
	}
	if len(chunk) > 0 {
		chunks = append(chunks, chunk)
	}

	var chunkFiles []string
	for _, c := range chunks {
		tempFile, err := sortAndWriteChunk(c)
		if err != nil {
			return err
		}
		chunkFiles = append(chunkFiles, tempFile)
	}

	err = PolyPhaseMergeSort(chunkFiles, outputFile)
	if err != nil {
		return err
	}

	for _, f := range chunkFiles {
		os.Remove(f)
	}

	return nil
}

// Function to perform sorting and writing of chunks to temporary files
func sortAndWriteChunk(chunk []int) (string, error) {
	sort.Ints(chunk)
	tempFile, err := os.CreateTemp("", "temp_chunk_*.txt")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	writer := bufio.NewWriter(tempFile)
	for _, item := range chunk {
		_, err := writer.WriteString(strconv.Itoa(item) + "\n")
		if err != nil {
			return "", err
		}
	}
	writer.Flush()

	return tempFile.Name(), nil
}

func main() {
	err := externalSort("input.txt", "output.txt")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("External sorting completed successfully.")
	}
}
