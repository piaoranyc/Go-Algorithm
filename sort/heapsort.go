package sort

func headify(arr []int, start int, end int) {
	left, right := start*2+1, start*2+2
	largest := start
	if left < end && arr[left] > arr[largest] {
		largest = left
	}
	if right < end && arr[right] > arr[largest] {
		largest = right
	}
	if largest != start {
		arr[largest], arr[start] = arr[start], arr[largest]
		headify(arr, largest, end)
	}

}

func heapsort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		headify(arr, i, len(arr))
	}
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		headify(arr, 0, i)
	}
}
