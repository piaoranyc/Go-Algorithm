package search

func binary_search(arr []int, target int) int {
	l, r := 0, len(arr)-1
	var mid int
	for l <= r {
		mid = l + (r-l)/2
		if arr[mid] > target {
			r = mid - 1
		} else if arr[mid] == target {
			return mid
		} else {
			l = mid + 1
		}
	}
	return -1
}
