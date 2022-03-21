package twoSum

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for i, num := range nums {
		comp := target - num
		if iComp, exists := indexMap[comp]; exists {
			return []int{iComp, i}
		}
		indexMap[num] = i
	}
	return nil
}
