package medianoftwosortedarray

import "sort"

func findMedianSortedArray(nums1 []int, nums2 []int) float64 {
	tmp := []int{}

	tmp = append(tmp, nums1...)
	tmp = append(tmp, nums2...)

	var result float64
	if len(tmp)%2 == 0 {
		sortedArray(tmp)
		result = float64(tmp[(len(tmp)/2)-1]+tmp[len(tmp)/2]) / 2
	} else {
		sortedArray(tmp)
		result = float64(tmp[len(tmp)/2])
	}

	return result
}

func sortedArray(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}
