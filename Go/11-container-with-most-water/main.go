package containerwithmostwater

import (
	"fmt"
	"time"
)

/*
** Optimized using Two Pointer Technique
** => O(n)
 */
func maxArea2(height []int) int {
	now := time.Now()
	defer func() {
		fmt.Println("Time Execution :", time.Since(now))
	}()

	result := 0

	left, right := 0, len(height)-1
	for left < right {
		// bandingkan dan ambil nilai terkecil dari height[left] dan height[right]
		tmp := min(height[left], height[right])

		// kalikan dengan jarak/step
		r := tmp * (right - left)

		// jika hasil perkalian lebih besar dari result, ganti nilai result
		if r > result {
			result = r
		}

		// jika value left lebih kecil dari right, increment left, selain itu decrement right
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
** Not Optimized - Brute Force Technique
** => O(n^2)
 */
// func maxArea(height []int) int {
// 	now := time.Now()
// 	defer func() {
// 		fmt.Println("Time Execution :", time.Since(now))
// 	}()

// 	tmp := 0
// 	r := 0

// 	step := 0
// 	for i, v := range height {
// 		for x := i; x < len(height); x++ {
// 			if v >= height[x] {
// 				tmp = height[x] * step
// 			} else {
// 				tmp = v * step
// 			}

// 			if tmp > r {
// 				r = tmp
// 			}

// 			step++
// 		}

// 		tmp = 0
// 		step = 0
// 	}

// 	return r
// }
