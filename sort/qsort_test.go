package sort

import (
	"fmt"
	"testing"
)

func TestQsort(t *testing.T) {
	arr := []int{4, 7, 8, 9, 1, 1, 4, 5, 5, 3, 11, 23, 54, 93, 65, 33, 43}
	//partition(arr, 0, len(arr)-1)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
