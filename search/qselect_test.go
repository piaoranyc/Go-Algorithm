package search

import (
	"fmt"
	"sort"
	"testing"
)

func TestQuickSelect(t *testing.T) {
	arr := []int{-44, 88, 44, 33, -1, 1, 2, 55, 3, 99, 4, 5, 6, 7, 8, 9}
	//partition(arr, 0, len(arr)-1)
	fmt.Println(arr)
	index := qselect(arr, 7, 0, len(arr)-1)
	println("index:", index, "arr[index]:", arr[index])
	fmt.Println(arr)

	sort.Ints(arr)
	fmt.Println(arr)
}
