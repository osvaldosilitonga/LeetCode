package twoSum

func TwoSum(nums []int, target int) []int {
	tmp := map[int]int{}

	for i, val := range nums {
		if mapIndex, found := tmp[target-val]; found {
			return []int{i, mapIndex}
		}
		tmp[val] = i
	}
	return nil
}
