package search

func partition(arr []int, left int, right int) int {
	key := left
	for left < right {
		for left < right && arr[key] <= arr[right] {
			right--
		}
		for left < right && arr[key] >= arr[left] {
			left++
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[left], arr[key] = arr[key], arr[left]
	key = left
	return key
}

func qselect(arr []int, k int, left int, right int) int {
	index := partition(arr, left, right)
	//fmt.Println(index)

	if index < k {
		return qselect(arr, k, index+1, right)
		//left = index + 1
	} else if index > k {
		return qselect(arr, k, left, index-1)
		//right = index - 1
	}

	return index
}
