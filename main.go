package main

import (
	"algorithms/sorting"
	"fmt"
)

func main() {
	arr := sorting.GenArr(10)
	fmt.Println("before sorting", arr)
	fmt.Println("after sorting", sorting.SelectionSort(arr))
}
