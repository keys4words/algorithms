package main

import (
	datastructures "algorithms/dataStructures"
	"fmt"
)

func main() {
	// arr := sorting.GenArr(10)
	// fmt.Println("before sorting", arr)
	// fmt.Println("after sorting", sorting.SelectionSort(arr))
	// fmt.Println("after sorting", sorting.BubbleSort(arr))
	// fmt.Println("after sorting", sorting.MergeSort(arr))

	//======================
	// datastructures.ArrayAddressing()
	var stack datastructures.Stack
	stack.Push("first string")
	stack.Push("second string")
	stack.Push("third string")
	for len(stack) > 0 {
		el, err := stack.Pop()
		if err == true {
			fmt.Println(el)
		}
	}
}
