package sort

func quickSort(arr []int, left int, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
}

//	func partition(arr []int, left int, right int) int {
//		pivot := arr[right]
//		i := left - 1
//		for j := left; j < right; j++ {
//			if arr[j] < pivot {
//				i++
//				arr[i], arr[j] = arr[j], arr[i]
//			}
//			fmt.Println(arr)
//		}
//		arr[i+1], arr[right] = arr[right], arr[i+1]
//		fmt.Println(arr)
//		return i + 1
//	}
func partition(arr []int, left int, right int) int {
	key := left

	for left < right {
		for left < right && arr[right] >= arr[key] {
			right--
		}
		for left < right && arr[left] <= arr[key] {
			left++
		}

		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[key], arr[left] = arr[left], arr[key]
	key = left

	return left
}
