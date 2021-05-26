package utils

func TwoSum(nums []int, target int) []int {
	mapObj := make(map[int]int)
	var result []int

	for i := 0; i < len(nums); i++ {
		if valueIndex, ok := mapObj[target - nums[i]]; ok {
			result = []int{valueIndex, i}
		}
		mapObj[nums[i]] = i
	}

	return result
}

