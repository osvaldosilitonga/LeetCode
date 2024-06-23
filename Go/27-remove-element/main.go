package removeelement

func removeElement(nums []int, val int) ([]int, int) {
	for i := 0; ; {
		if i == len(nums) {
			return nums, len(nums)
		}

		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}

		i++
	}
}
