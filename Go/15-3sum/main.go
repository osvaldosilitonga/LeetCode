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

// optimal solution. time complexity: O(n^2), space complexity: O(1)
func threeSum2(nums []int) [][]int {
	n := len(nums)

	sort.Ints(nums)

	result := [][]int{}
	for i := 0; i < n-2; i++ {
		// skip all duplicate from left
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				r := []int{nums[i], nums[j], nums[k]}
				result = append(result, r)

				k--

				// Skip all duplicates from right
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				//  decrement k to reduce sum value
				k--
			} else {
				// increment j to increase sum value
				j++
			}
		}
	}

	return result
}
