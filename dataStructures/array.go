package datastructures

import "fmt"

func ArrayAddressing() {
	arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range arr {
		fmt.Println("Element", i, "has value", v, " and memory location address", &arr[i])
	}
}
