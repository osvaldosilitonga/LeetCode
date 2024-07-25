package threesum

import (
	"sort"
)

// Given an integer array nums, return triplets [nums[i], nums[j], nums[k]]
// such that i != j, i != k, and j != k
// and nums[i] + nums[j] + nums[k] == 0
func threeSum(nums []int) [][]int {
	n := len(nums)

	numsOfTriplets := make(map[[3]int][]int)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				triplets := [3]int{nums[i], nums[j], nums[k]}
				sort.Ints(triplets[:])

				if nums[i]+nums[j]+nums[k] == 0 {
					numsOfTriplets[triplets] = triplets[:]
				}
			}
		}
	}

	result := [][]int{}
	for _, v := range numsOfTriplets {
		result = append(result, v)
	}

	return result
}
