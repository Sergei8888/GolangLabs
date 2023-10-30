package main

import (
	"reflect"
	"testing"
)

func TestCombSort(t *testing.T) {
	// Given
	testSlice1 := []int{1, 74, 23, 423, 2, 34, 2}
	testSlice2 := []float64{1.22, 5.63, 1, 8.234, 1123.1, 5.6}

	// When
	testSlice1 = CombSort(testSlice1)
	testSlice2 = CombSort(testSlice2)

	// Then
	if !reflect.DeepEqual(testSlice1, []int{1, 2, 2, 23, 34, 74, 423}) {
		t.Fatal("Slices are not equal")
	}
	if !reflect.DeepEqual(testSlice2, []float64{1, 1.22, 5.6, 5.63, 8.234, 1123.1}) {
		t.Fatal("Slices are not equal")
	}
}

func TestInsertionSort(t *testing.T) {
	// Given
	testSlice1 := []int{1, 74, 23, 423, 2, 34, 2}
	testSlice2 := []float64{1.22, 5.63, 1, 8.234, 1123.1, 5.6}

	// When
	testSlice1 = InsertionSort(testSlice1)
	testSlice2 = InsertionSort(testSlice2)

	// Then
	if !reflect.DeepEqual(testSlice1, []int{1, 2, 2, 23, 34, 74, 423}) {
		t.Fatal("Slices are not equal")
	}
	if !reflect.DeepEqual(testSlice2, []float64{1, 1.22, 5.6, 5.63, 8.234, 1123.1}) {
		t.Fatal("Slices are not equal")
	}
}

func TestSelectionSort(t *testing.T) {
	// Given
	testSlice1 := []int{1, 74, 23, 423, 2, 34, 2}
	testSlice2 := []float64{1.22, 5.63, 1, 8.234, 1123.1, 5.6}

	// When
	testSlice1 = SelectionSort(testSlice1)
	testSlice2 = SelectionSort(testSlice2)

	// Then
	if !reflect.DeepEqual(testSlice1, []int{1, 2, 2, 23, 34, 74, 423}) {
		t.Fatal("Slices are not equal")
	}
	if !reflect.DeepEqual(testSlice2, []float64{1, 1.22, 5.6, 5.63, 8.234, 1123.1}) {
		t.Fatal("Slices are not equal")
	}
}
