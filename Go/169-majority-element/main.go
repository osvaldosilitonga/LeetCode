package majorityelement

func majorityElement(nums []int) int {
	tmp := map[int]int{}

	for _, v := range nums {
		tmp[v] += 1
	}

	var key int
	var value int
	for k, v := range tmp {
		if v > value {
			key = k
			value = v
		}
	}

	return key
}
