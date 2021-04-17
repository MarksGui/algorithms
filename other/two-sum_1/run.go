package two_sum_1

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if p, ok := m[target-num]; ok {
			return []int{p, i}
		}
		m[num] = i
	}
	return nil
}
