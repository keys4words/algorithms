package main

import (
	"algorithms/sorting"
	"fmt"
)

func main() {
	arr := []int{8, 2, 9, 0, 4, 1}
	fmt.Println("before sorting", arr)
	// sorting.BubbleSort(arr)
	sorting.InsertionSort(arr)
	fmt.Println("after sorting", arr)
}
