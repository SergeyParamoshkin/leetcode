package twosum

func TwoSum(nums []int, target int) []int {
	a := []int{}
	m := make(map[int]int, len(nums))
	for index, value := range nums {
		if idx, ok := m[target-value]; ok {
			a = append(a, idx, index)
		}
		m[value] = index
	}

	return a
}
