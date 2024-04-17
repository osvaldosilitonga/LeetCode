package twoSum

func TwoSum(nums []int, target int) []int {
	tmp := map[int]int{}

	for i, val := range nums {
		if tmpIndex, found := tmp[target-val]; found {
			return []int{i, tmpIndex}
		}
		tmp[val] = i
	}
	return nil
}
