package sorting

func BubbleSort(arr []int) {
	isDone := false
	for !isDone {
		isDone = true
		i := 0
		for i < len(arr)-1 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isDone = false
			}
			i++
		}
	}
}

func InsertionSort(arr []int) {
	i := 1
	for i < len(arr) {
		j := i
		for j >= 1 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
		i++
	}
}
