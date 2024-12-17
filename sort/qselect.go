package sort

func qselect(arr []int, k int, left int, right int) int {
	index := partition(arr, left, right)
	//fmt.Println(index)

	if index < k {
		return qselect(arr, k, index+1, right)
	} else if index > k {

		return qselect(arr, k, left, index-1)
		//right = index - 1
	}

	return index
}
