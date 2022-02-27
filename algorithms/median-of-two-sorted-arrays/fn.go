package median_of_two_sorted_arrays

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numbers := append(nums1, nums2...)
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })
	fmt.Println(numbers)
	if len(numbers) % 2 != 0 {
		n := numbers[len(numbers) / 2]
		return float64(n)
	}
	m := len(numbers) / 2
	return (float64(numbers[m-1]) + float64(numbers[m])) / 2.0
}
