package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	// Expected: 2
	input := []int{-1, 2, 1, -4}
	result := threeSumClosest(input, 1)
	fmt.Printf("Input : %v, Target : %v, Result : %v", input, 1, result)

	fmt.Println()

	// Expected: 0
	input2 := []int{0, 0, 0}
	result2 := threeSumClosest(input2, 1)
	fmt.Printf("Input : %v, Target : %v, Result : %v", input2, 1, result2)

	fmt.Println()

	// Expected: 2
	input3 := []int{1, 1, 1, 0}
	result3 := threeSumClosest(input3, -100)
	fmt.Printf("Input : %v, Target : %v, Result : %v", input3, -100, result3)

	fmt.Println()

	// Expected: 3
	input4 := []int{0, 1, 2}
	result4 := threeSumClosest(input4, 3)
	fmt.Printf("Input : %v, Target : %v, Result : %v", input4, 3, result4)

	fmt.Println()

	// Expected: -2
	input5 := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	result5 := threeSumClosest(input5, -2)
	fmt.Printf("Input : %v, Target : %v, Result : %v", input5, -2, result5)

	fmt.Println()

	// Expected: 300
	input6 := []int{0, 3, 97, 102, 200}
	result6 := threeSumClosest(input6, 300)
	fmt.Printf("Input : %v, Target : %v, Result : %v", input6, 300, result6)

	fmt.Println()

	// Expected: 2
	input7 := []int{1, 1, 1, 0}
	result7 := threeSumClosest(input7, -100)
	fmt.Printf("Input : %v, Target : %v, Result : %v", input7, -100, result7)
}

// ATTEMPT 2
func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)

	numsLen := len(nums)
	var tmp int
	var result int
	distance := math.MaxInt32

	for i := 0; i < numsLen-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		low, high := i+1, numsLen-1

		for low < high {
			tmp = nums[low] + nums[i] + nums[high]
			dst := math.Abs(float64(target - tmp))
			if dst < float64(distance) {
				distance = int(dst)
				result = tmp
			}

			if tmp == target {
				return tmp
			} else if tmp < target {
				low++
			} else if tmp > target {
				high--
			}
		}
	}

	return result
}

// ATTEMPT 1
// func threeSumClosest(nums []int, target int) int {
// 	var result int
// 	lenNums := len(nums)
// 	sum := 0
// 	count := 1

// 	for i := 0; i < lenNums; i++ {
// 		sum += nums[i]

// 		if result == target {
// 			return result
// 		}

// 		if count == 3 {
// 			if target >= 0 && (target-sum) <= (target-result) {
// 				result = sum
// 			} else if target < 0 && (target-sum) <= (result-target) {
// 				result = sum
// 			}

// 			sum -= nums[i-2]
// 		}

// 		// fmt.Println("=====================================")
// 		// fmt.Println("Count =", count)
// 		// fmt.Println("Result =", result)
// 		// fmt.Println("Sum =", sum)
// 		// fmt.Println("Target - Sum =", target-sum)
// 		// fmt.Println("Target - Result =", target-result)
// 		// fmt.Println("sum < target :", (target-sum) < (target-result))

// 		if count < 3 {
// 			count++
// 		}
// 	}

// 	return result
// }
